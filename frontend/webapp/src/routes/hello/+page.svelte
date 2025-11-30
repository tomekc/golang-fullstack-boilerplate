<script lang="ts">
    import { fetchCurrentTime } from '$lib/services/hello';
    import { Button, Card, CardBody, CardFooter, CardHeader } from '@sveltestrap/sveltestrap';

    let currentTime = $state<string>('');
    let loading = $state<boolean>(false);
    let error = $state<string>('');

    async function handleGetTime() {
        loading = true;
        error = '';
        currentTime = '';

        try {
            const response = await fetchCurrentTime();
            currentTime = response.time;
        } catch (err) {
            error = err instanceof Error ? err.message : 'Unknown error';
        } finally {
            loading = false;
        }
    }
</script>

<h1>Hello World</h1>
<p class="text-muted">Interactive backend call example</p>

<div class="mt-4">
    <Card class="shadow-sm" style="max-width: 600px;">
        <CardHeader>
            <h5 class="mb-0">Backend Call</h5>
        </CardHeader>
        <CardBody>
            {#if loading}
                <div class="d-flex align-items-center">
                    <div class="spinner-border spinner-border-sm text-primary me-2" role="status">
                        <span class="visually-hidden">Loading...</span>
                    </div>
                    <span>Fetching time from server...</span>
                </div>
            {:else if error}
                <div class="alert alert-danger mb-0" role="alert">
                    <strong>Error:</strong> {error}
                </div>
            {:else if currentTime}
                <div class="alert alert-success mb-0" role="alert">
                    <strong>Server time:</strong> {currentTime}
                </div>
            {:else}
                <p class="text-muted mb-0">
                    Click the button below to fetch the current time from the backend server.
                </p>
            {/if}
        </CardBody>
        <CardFooter>
            <Button color="primary" onclick={handleGetTime} disabled={loading}>
                {loading ? 'Loading...' : 'Get the time'}
            </Button>
        </CardFooter>
    </Card>
</div>
