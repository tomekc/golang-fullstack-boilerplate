<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        children?: Snippet;
        userInfo?: string;
        width?: string;
    }

    let { children, userInfo, width = '225px' }: Props = $props();
</script>

<aside class="menu sb-sidenav" style="flex-basis: {width}; flex-shrink: 0;">
    <div class="sb-sidenav-inner is-flex is-flex-direction-column" style="height: 100%;">
        <div class="sb-sidenav-menu is-flex-grow-1" style="overflow-y: auto;">
            {@render children?.()}
        </div>

        {#if userInfo}
            <div class="sb-sidenav-footer">
                <p class="is-size-7 has-text-grey-light">Logged in as:</p>
                <p class="has-text-white">{userInfo}</p>
            </div>
        {/if}
    </div>
</aside>

<style>
    .sb-sidenav {
        background-color: #212529;
        color: rgba(255, 255, 255, 0.5);
        transition: transform 0.15s ease-in-out;
    }

    .sb-sidenav-footer {
        padding: 0.75rem 1rem;
        background-color: rgba(0, 0, 0, 0.1);
    }

    :global(.sb-sidenav .menu-label) {
        padding: 1.75rem 1rem 0.75rem;
        font-size: 0.75rem;
        font-weight: 800;
        text-transform: uppercase;
        letter-spacing: 0.05em;
        color: rgba(255, 255, 255, 0.25);
        margin-bottom: 0;
    }

    :global(.sb-sidenav .menu-list a) {
        color: rgba(255, 255, 255, 0.5);
        background-color: transparent;
        padding: 0.75rem 1rem;
        border-radius: 0;
        transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out;
    }

    :global(.sb-sidenav .menu-list a:hover) {
        color: #fff;
        background-color: rgba(255, 255, 255, 0.075);
    }

    :global(.sb-sidenav .menu-list a.is-active) {
        color: #fff;
        background-color: rgba(255, 255, 255, 0.1);
    }

    @media (max-width: 768px) {
        .sb-sidenav {
            transform: translateX(-225px);
            position: fixed;
            top: 56px;
            bottom: 0;
            z-index: 1038;
        }
    }
</style>
