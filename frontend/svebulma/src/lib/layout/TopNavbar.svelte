<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        appName?: string;
        version?: string;
        brand?: Snippet;
        end?: Snippet;
    }

    let { appName = 'App', version, brand, end }: Props = $props();
</script>

<nav class="navbar is-dark sb-topnav">
    <div class="navbar-brand">
        {#if brand}
            {@render brand()}
        {:else}
            <a class="navbar-item has-text-weight-bold" href="/">
                {appName}
            </a>
        {/if}
    </div>
    <div class="navbar-menu is-active">
        <div class="navbar-end">
            {#if end}
                {@render end()}
            {/if}
            {#if version}
                <div class="navbar-item">
                    <span class="has-text-grey-light is-size-7">v{version}</span>
                </div>
            {/if}
        </div>
    </div>
</nav>

<style>
    .sb-topnav {
        z-index: 1039;
        height: 56px;
        flex-shrink: 0;
    }
</style>
