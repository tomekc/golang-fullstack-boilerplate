<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        href: string;
        icon?: string;
        active?: boolean;
        hasUpdateMark?: boolean;
        children?: Snippet;
    }

    let { href, icon, active = false, hasUpdateMark = false, children }: Props = $props();
</script>

<li>
    <a {href} class:is-active={active} class:has-icon={!!icon}>
        {#if icon}
            <span class="icon" class:has-update-mark={hasUpdateMark}><i class="mdi mdi-{icon}"></i></span>
        {/if}
        <span class="menu-item-label">{@render children?.()}</span>
    </a>
</li>

<style>
    li {
        list-style: none;
    }
</style>
