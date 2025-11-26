<script lang="ts">
	import { onMount } from 'svelte';
	import { fetchCurrentTime } from '../lib/services/hello';
    import {Button, Col, Row} from "@sveltestrap/sveltestrap";

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

<Row>
    <h1>Welcome to Golang-SvelteKit Boilerplate</h1>
</Row>
<Row class="pt-4">
    <Col md="2">
        Get time from backend:
    </Col>
    <Col md="9">
        {#if loading}
            <p>Loading...</p>
        {:else if error}
            <p>Error: {error}</p>
        {:else}
            <p>Now is {currentTime}</p>
        {/if}
    </Col>
</Row>

<Row class="bg-secondary-subtle">
    <div class="p-3">
        Made with <a href="https://svelte.dev/docs/kit"><b>Svelte</b></a>
            and <a href="https://sveltestrap.js.org/?path=/docs/sveltestrap-overview--docs"><b>Sveltestrap</b></a> (Bootstrap CSS components).
    </div>
</Row>
