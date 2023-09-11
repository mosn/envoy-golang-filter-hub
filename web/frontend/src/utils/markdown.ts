// import { unified } from 'unified'
// import remarkParse from 'remark-parse'
// import remarkGfm from 'remark-gfm'
// import remarkRehype from 'remark-rehype'
// import rehypeExternalLinks from 'rehype-external-links'
// import rehypeStringify from 'rehype-stringify'
// import rehypeSlug from 'rehype-slug'
// import rehypeAutolinkHeadings from 'rehype-autolink-headings'
// import rehypeRaw from 'rehype-raw'
// import rehypeSanitize from 'rehype-sanitize'
// import sanitizeScheme from "./sanitize-schema"

// const processor = unified()
//   .use(remarkParse)
//   .use(remarkGfm)
//   .use(remarkRehype, { allowDangerousHtml: true })

// processor
//   .use(rehypeStringify)
//   .use(rehypeRaw)
//   .use(rehypeSlug)
//   .use(rehypeAutolinkHeadings)
//   .use(rehypeSanitize, sanitizeScheme)
//   .use(rehypeExternalLinks, { rel: ['nofollow', 'noopener'], target: '_blank' })

// export const markdownToHtml = async (markdown: string) => {
//   const result = await processor.process(markdown)
//   return result.toString()
// }
