<script lang="ts">
    import { fetchCurrentTime } from '$lib/services/hello';

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

<h1 class="title">Hello World</h1>
<p class="subtitle has-text-grey">Interactive backend call example</p>

<div class="mt-4" style="max-width: 600px;">
    <div class="card">
        <div class="card-header">
            <p class="card-header-title">Backend Call</p>
        </div>
        <div class="card-content">
            {#if loading}
                <div class="is-flex is-align-items-center">
                    <span class="icon has-text-primary mr-2">
                        <i class="fas fa-spinner fa-spin"></i>
                    </span>
                    <span>Fetching time from server...</span>
                </div>
            {:else if error}
                <div class="notification is-danger">
                    <strong>Error:</strong> {error}
                </div>
            {:else if currentTime}
                <div class="notification is-success">
                    <strong>Server time:</strong> {currentTime}
                </div>
            {:else}
                <p class="has-text-grey">
                    Click the button below to fetch the current time from the backend server.
                </p>
            {/if}
        </div>
        <div class="card-footer">
            <div class="card-footer-item">
                <button class="button is-primary" onclick={handleGetTime} disabled={loading}>
                    {loading ? 'Loading...' : 'Get the time'}
                </button>
            </div>
        </div>
    </div>
</div>
