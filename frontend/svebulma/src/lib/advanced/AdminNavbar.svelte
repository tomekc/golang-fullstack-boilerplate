<script lang="ts">
    import type { Snippet } from 'svelte';
    import NavbarSearch from './NavbarSearch.svelte';

    interface Props {
        searchPlaceholder?: string;
        navbarEnd?: Snippet;
    }

    let { searchPlaceholder = 'Search everywhere...', navbarEnd }: Props = $props();
</script>

<nav id="navbar-main" class="navbar is-fixed-top">
    <div class="navbar-brand">
        <a class="navbar-item is-hidden-desktop jb-aside-mobile-toggle" href={undefined} aria-label="Toggle sidebar">
            <span class="icon"><i class="mdi mdi-forwardburger mdi-24px"></i></span>
        </a>
        <NavbarSearch placeholder={searchPlaceholder} />
    </div>
    <div class="navbar-brand is-right">
        <a class="navbar-item is-hidden-desktop jb-navbar-menu-toggle" href={undefined} data-target="navbar-menu" aria-label="Toggle menu">
            <span class="icon"><i class="mdi mdi-dots-vertical"></i></span>
        </a>
    </div>
    <div class="navbar-menu fadeIn animated faster" id="navbar-menu">
        <div class="navbar-end">
            {#if navbarEnd}
                {@render navbarEnd()}
            {/if}
        </div>
    </div>
</nav>
