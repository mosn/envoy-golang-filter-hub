import os
import subprocess
import requests
import json

GITHUB_TOKEN = os.getenv("GITHUB_TOKEN")
GITHUB_REPOSITORY = os.getenv("GITHUB_REPOSITORY")
API_URL = "https://api.github.com/markdown"


def render_markdown_to_html(plugin_path):
    readme_path = os.path.join(plugin_path, "readme.md")
    if not os.path.exists(readme_path):
        print(f"Warning: No readme.md found in {plugin_path}")
        return

    with open(readme_path, "r", encoding="utf-8") as f:
        markdown_content = f.read()

    payload = json.dumps({
        "text": markdown_content,
        "mode": "gfm",
        "context": GITHUB_REPOSITORY
    })
    headers = {
        'Authorization': f'token {GITHUB_TOKEN}',
        'Content-Type': 'application/json',
        'Accept': 'application/vnd.github.v3+json'
    }
    response = requests.post(API_URL, headers=headers, data=payload)

    if response.status_code == 200:
        html_content = response.text
        print(f"Rendered markdown to html for {plugin_path}")
        html_path = os.path.join(plugin_path, "readme.html")
        with open(html_path, "w", encoding="utf-8") as f:
            f.write(html_content)
    else:
        print(f"Error: Could not render markdown for {plugin_path}")


def commit_and_push(plugin_path):
    subprocess.run(["git", "add", f"{plugin_path}/readme.html"])
    subprocess.run(["git", "commit", "-m", f"Update readme.html for {plugin_path}"])
    subprocess.run(["git", "push", f"https://x-access-token:{GITHUB_TOKEN}@github.com/{GITHUB_REPOSITORY}.git"])


def main():
    plugins_path = "./plugins"  # 这应当是插件目录的路径
    changed_files = os.getenv('CHANGED_FILES').split() if os.getenv('CHANGED_FILES') else []
    changed_plugins = set(file.split('/')[1] for file in changed_files if file.startswith('plugins/'))

    for plugin_name in changed_plugins:
        plugin_path = os.path.join(plugins_path, plugin_name)
        render_markdown_to_html(plugin_path)
        commit_and_push(plugin_path)


if __name__ == "__main__":
    main()
