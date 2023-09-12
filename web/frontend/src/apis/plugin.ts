import { baseUrl } from '@/utils/env'

export interface PluginItem {
  category: string
  description: string
  name: string
  path_name: string
  version: string
}

export const getPluginList = async (): Promise<PluginItem[] | false> => {
  return await fetch(baseUrl + 'pluginList')
    .then((res) => res.json())
    .then((res) => {
      if (res.code !== 200) throw new Error(res.message)

      return res.data.plugins as PluginItem[]
    })
    .catch((err) => {
      console.error(err)
      return false
    })
}

export type PluginData = PluginItem & {
  versions: PluginVersionItem[]
  overview?: string
  config?: string
  changelog?: string
}

export interface PluginVersionItem {
  version: string
  created_at: string
  commit_hash: string
  commit_url: string
  downloads: PluginDownloadItem[]
}

export interface PluginDownloadItem {
  type: string
  url: string
}

export const getPlugin = async (path: string): Promise<PluginData | false> => {
  return await fetch(baseUrl + 'plugin/' + path)
    .then((res) => res.json())
    .then((res) => {
      if (res.code !== 200) throw new Error(res.message)

      return res.data as PluginData
    })
    .catch((err) => {
      console.error(err)
      return false
    })
}
