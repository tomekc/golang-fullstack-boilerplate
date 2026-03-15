<script lang="ts">
    import { fetchCurrentTime } from '$lib/services/hello';
    import { baseURL } from '$lib/config';
    import { TitleBar, HeroBar, Card, Button, Notification, Icon, Table } from 'svebulma';

    let currentTime = $state<string>('');
    let loading = $state<boolean>(false);
    let error = $state<string>('');

    let echoBody = $state<Record<string, unknown> | null>(null);
    let echoLoading = $state<boolean>(false);
    let echoError = $state<string>('');

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

    async function handleEcho() {
        echoLoading = true;
        echoError = '';
        echoBody = null;

        const browserInfo = {
            userAgent: navigator.userAgent,
            language: navigator.language,
            platform: navigator.platform,
            screenWidth: screen.width,
            screenHeight: screen.height,
            colorDepth: screen.colorDepth,
            timestamp: new Date().toISOString()
        };

        try {
            const res = await fetch(`${baseURL()}/api/echo`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(browserInfo)
            });
            if (!res.ok) throw new Error(`Request failed: ${res.statusText}`);
            const data = await res.json();
            echoBody = data.body;
        } catch (err) {
            echoError = err instanceof Error ? err.message : 'Unknown error';
        } finally {
            echoLoading = false;
        }
    }
</script>

<TitleBar breadcrumbs={['Admin', 'Hello API']} />
<HeroBar title="Hello API" />

<section class="section is-main-section">
    <div style="max-width: 600px;">
        <Card title="Server Call" icon="clock-outline">
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
                    Click the button below to fetch the current time from the server.
                </p>
            {/if}
        </Card>
    </div>

    <div style="max-width: 600px; margin-top: 1.5rem;">
        <Card title="Echo POST" icon="send" hasTable={!!echoBody}>
            {#snippet footer()}
                <div class="card-footer-item">
                    <Button color="info" loading={echoLoading} disabled={echoLoading} onclick={handleEcho}>
                        {echoLoading ? 'Sending...' : 'Send POST'}
                    </Button>
                </div>
            {/snippet}

            {#if echoLoading}
                <div class="is-flex is-align-items-center">
                    <Icon name="loading" color="info" spin />&nbsp;
                    <span>Sending browser info...</span>
                </div>
            {:else if echoError}
                <Notification color="danger"><strong>Error:</strong> {echoError}</Notification>
            {:else if echoBody}
                <Table>
                    {#snippet head()}
                        <tr>
                            <th>Key</th>
                            <th>Value</th>
                        </tr>
                    {/snippet}
                    {#each Object.entries(echoBody) as [key, value]}
                        <tr>
                            <td>{key}</td>
                            <td>{value}</td>
                        </tr>
                    {/each}
                </Table>
            {:else}
                <p class="has-text-grey">
                    Click the button below to send browser info to the echo endpoint.
                </p>
            {/if}
        </Card>
    </div>
</section>
