<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        striped?: boolean;
        hoverable?: boolean;
        bordered?: boolean;
        fullwidth?: boolean;
        narrow?: boolean;
        scrollable?: boolean;
        head?: Snippet;
        foot?: Snippet;
        children?: Snippet;
    }

    let {
        striped = false,
        hoverable = true,
        bordered = false,
        fullwidth = true,
        narrow = false,
        scrollable = false,
        head,
        foot,
        children,
    }: Props = $props();

    const classes = $derived([
        'table',
        striped ? 'is-striped' : '',
        hoverable ? 'is-hoverable' : '',
        bordered ? 'is-bordered' : '',
        fullwidth ? 'is-fullwidth' : '',
        narrow ? 'is-narrow' : '',
    ].filter(Boolean).join(' '));
</script>

{#if scrollable}
    <div class="table-container">
        <table class={classes}>
            {#if head}<thead>{@render head()}</thead>{/if}
            {#if foot}<tfoot>{@render foot()}</tfoot>{/if}
            <tbody>{@render children?.()}</tbody>
        </table>
    </div>
{:else}
    <table class={classes}>
        {#if head}<thead>{@render head()}</thead>{/if}
        {#if foot}<tfoot>{@render foot()}</tfoot>{/if}
        <tbody>{@render children?.()}</tbody>
    </table>
{/if}
