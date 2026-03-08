<script lang="ts">
    import { fetchCurrentTime } from '$lib/services/hello';
    import { TitleBar, HeroBar, Card, Button, Notification, Icon } from 'svebulma';

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

<TitleBar breadcrumbs={['Admin', 'Hello API']} />
<HeroBar title="Hello API" />

<section class="section is-main-section">
    <div style="max-width: 600px;">
        <Card title="Backend Call" icon="clock-outline">
            {#snippet footer()}
                <div class="card-footer-item">
                    <Button color="primary" loading={loading} disabled={loading} onclick={handleGetTime}>
                        {loading ? 'Loading...' : 'Get the time'}
                    </Button>
                </div>
            {/snippet}

            {#if loading}
                <div class="is-flex is-align-items-center">
                    <Icon name="loading" color="primary" spin />&nbsp;
                    <span>Fetching time from server...</span>
                </div>
            {:else if error}
                <Notification color="danger"><strong>Error:</strong> {error}</Notification>
            {:else if currentTime}
                <Notification color="success"><strong>Server time:</strong> {currentTime}</Notification>
            {:else}
                <p class="has-text-grey">
                    Click the button below to fetch the current time from the backend server.
                </p>
            {/if}
        </Card>
    </div>
</section>
