<script lang="ts">
    import { page } from '$app/state';

    interface Props {
        href: string;
        icon?: string;
    }

    let { href, icon, children }: Props = $props();

    const isActive = page.url.pathname === href;
</script>

<li>
    <a {href} class="is-flex is-align-items-center" class:is-active={isActive}>
        {#if icon}
            <span class="icon is-small mr-2">
                <i class="fas fa-{icon}"></i>
            </span>
        {/if}
        <span>{@render children?.()}</span>
    </a>
</li>

<style>
    li {
        list-style: none;
    }

    a {
        display: flex;
        align-items: center;
        padding: 0.75rem 1rem;
        color: rgba(255, 255, 255, 0.5);
        text-decoration: none;
        transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out;
        border-radius: 0;
    }

    a:hover {
        color: #fff;
        background-color: rgba(255, 255, 255, 0.075);
    }

    a.is-active {
        color: #fff !important;
        background-color: rgba(255, 255, 255, 0.1);
    }
</style>
