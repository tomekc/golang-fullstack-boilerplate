<script lang="ts">
    import { page } from '$app/state';
    import { NavLink } from '@sveltestrap/sveltestrap';

    interface Props {
        href: string;
        icon?: string;
        [key: string]: any;
    }

    let { href, icon, children, ...restProps }: Props = $props();

    console.log(`Href=${href} Page=${page.url.pathname}`);

    const isActive = page.url.pathname === href;
    let klas = isActive ? 'active' : '';
</script>

<NavLink {href} class="nav-link {klas}" {...restProps}>
    {#if icon}
        <div class="sb-nav-link-icon">
            <i class="fas fa-{icon}"></i>
        </div>
    {/if}
    {@render children?.()}
</NavLink>

<style>
    :global(.sb-sidenav .nav-link) {
        display: flex;
        align-items: center;
        padding: 0.75rem 1rem;
        color: rgba(255, 255, 255, 0.5);
        text-decoration: none;
        transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out;
    }

    :global(.sb-sidenav .nav-link:hover) {
        color: #fff;
        background-color: rgba(255, 255, 255, 0.075);
    }

    :global(.sb-sidenav .nav-link.active) {
        color: #fff !important;
        background-color: rgba(255, 255, 255, 0.1);
    }

    :global(.sb-nav-link-icon) {
        display: inline-flex;
        align-items: center;
        justify-content: center;
        width: 1rem;
        margin-right: 0.75rem;
        font-size: 0.9rem;
    }
</style>
