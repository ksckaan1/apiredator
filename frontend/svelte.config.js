import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),
	kit: {
		adapter: adapter(),
		alias: {
			"$components": "src/components",
			"$styles": "src/lib/style",
			"$types": "src/lib/types",
			"$assets": "src/lib/assets",
			"$utils": "src/lib/utils",
			"$stores": "src/store",

		}
	},
};

export default config;
