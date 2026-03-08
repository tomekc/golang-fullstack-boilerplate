<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        href: string;
        icon?: string;
        active?: boolean;
        children?: Snippet;
    }

    let { href, icon, active = false, children }: Props = $props();
</script>

<li>
    <a {href} class:is-active={active}>
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
</style>
