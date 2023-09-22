export const github_repo = import.meta.env.VITE_GITHUB_REPO
export const baseUrl = "https://ghproxy.com/https://raw.githubusercontent.com/" + github_repo + "/main/web/cache/"
export const isDev = import.meta.env.DEV