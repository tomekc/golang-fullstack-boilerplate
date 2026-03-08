<script lang="ts">
    import type { Snippet } from 'svelte';
    import TopNavbar from './TopNavbar.svelte';
    import Sidebar from './Sidebar.svelte';
    import MainContent from './MainContent.svelte';
    import PageFooter from './PageFooter.svelte';

    interface Props {
        appName?: string;
        version?: string;
        userInfo?: string;
        copyright?: string;
        sidebar?: Snippet;
        navEnd?: Snippet;
        footer?: Snippet;
        children?: Snippet;
    }

    let { appName, version, userInfo, copyright, sidebar, navEnd, footer, children }: Props = $props();
</script>

<div class="is-flex is-flex-direction-column" style="min-height: 100vh;">
    <TopNavbar {appName} {version} end={navEnd} />
    <div class="is-flex is-flex-grow-1">
        {#if sidebar}
            <Sidebar {userInfo}>
                {@render sidebar()}
            </Sidebar>
        {/if}
        <MainContent>
            {@render children?.()}
        </MainContent>
    </div>
    {#if footer}
        <PageFooter>
            {@render footer()}
        </PageFooter>
    {:else if copyright}
        <PageFooter {copyright} />
    {/if}
</div>
