<script lang="ts">
    interface Props {
        icon?: string;
        label?: string;
        href?: string;
        divider?: boolean;
    }

    let { icon = undefined, label = undefined, href = undefined, divider = false }: Props = $props();
</script>

{#if divider}
    <hr class="navbar-divider">
{:else}
    <a class="navbar-item" href={href}>
        {#if icon}
            <span class="icon"><i class="mdi mdi-{icon}"></i></span>
        {/if}
        <span>{label}</span>
    </a>
{/if}
