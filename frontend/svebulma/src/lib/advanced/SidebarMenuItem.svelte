<script lang="ts">
    interface Props {
        label: string;
        href?: string;
        icon?: string;
        active?: boolean;
        onclick?: () => void;
    }

    let { label, href = undefined, icon = undefined, active = false, onclick = undefined }: Props = $props();
</script>

<li>
    <a {href} class:has-icon={icon} class:is-active={active} {onclick}>
        {#if icon}
            <span class="icon"><i class="mdi mdi-{icon}"></i></span>
        {/if}
        <span class="menu-item-label">{label}</span>
    </a>
</li>
