<script lang="ts">
	import { onMount } from 'svelte';
	import { fetchCurrentTime } from '$lib/services/hello';

	let currentTime = $state<string>('');
	let loading = $state<boolean>(true);
	let error = $state<string>('');

	onMount(async () => {
		try {
			const response = await fetchCurrentTime();
			currentTime = response.time;
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
		} finally {
			loading = false;
		}
	});
</script>

<h1>Welcome to Golang-SvelteKit Boilerplate</h1>

{#if loading}
	<p>Loading...</p>
{:else if error}
	<p>Error: {error}</p>
{:else}
	<p>Now is {currentTime}</p>
{/if}

<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
