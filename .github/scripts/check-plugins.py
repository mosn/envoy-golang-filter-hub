import os
import sys
import yaml


def check_each_plugin(plugin_dir):
    print(plugin_dir)
    # 检查插件是否符合规范，如果有错误，输出错误信息并返回 False
    plugin_file = os.path.join(plugin_dir, 'metadata.yaml')
    # 如果文件不存在
    if not os.path.isfile(plugin_file):
        print(f"错误：插件 {plugin_dir} 缺少 metadata.yaml 文件")
        return False

    with open(plugin_file) as f:
        plugin = yaml.safe_load(f)
    if not plugin.get('name'):
        print(f"错误：插件 {plugin_dir} 的名称不能为空")
        return False
    if not plugin.get('version'):
        print(f"错误：插件 {plugin_dir} 的版本号不能为空")
        return False
    # 其他检查逻辑...
    return True


def main(plugins_dir):
    # 检查每个插件是否符合规范
    for plugin_name in os.listdir(plugins_dir):
        plugin_path = os.path.join(plugins_dir, plugin_name)
        if not os.path.isdir(plugin_path):
            continue
        if not check_each_plugin(plugin_path):
            sys.exit(1)  # 如果出现错误，退出脚本并返回非零退出码


if __name__ == '__main__':
    print("python start")
    main(os.path.abspath("./plugins"))
