<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        label: string;
        avatarUrl?: string;
        hasDivider?: boolean;
        children: Snippet;
    }

    let { label, avatarUrl = undefined, hasDivider = false, children }: Props = $props();
</script>

<div
    class="navbar-item has-dropdown has-dropdown-with-icons has-user-avatar is-hoverable"
    class:has-divider={hasDivider}
>
    <a class="navbar-link is-arrowless" href={undefined}>
        {#if avatarUrl}
            <div class="is-user-avatar">
                <img src={avatarUrl} alt={label}>
            </div>
        {/if}
        <div class="is-user-name"><span>{label}</span></div>
        <span class="icon"><i class="mdi mdi-chevron-down"></i></span>
    </a>
    <div class="navbar-dropdown">
        {@render children()}
    </div>
</div>
