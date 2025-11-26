import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	optimizeDeps: {
		include: ['@sveltestrap/sveltestrap']
	},
	ssr: {
		noExternal: ['@sveltestrap/sveltestrap']
	}
});
