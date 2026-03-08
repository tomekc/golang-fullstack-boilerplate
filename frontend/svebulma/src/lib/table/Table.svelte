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
        pagination?: Snippet;
        children?: Snippet;
    }

    let {
        striped = true,
        hoverable = true,
        bordered = false,
        fullwidth = true,
        narrow = false,
        scrollable = false,
        head,
        foot,
        pagination,
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

<div class="b-table">
    {#if scrollable}
        <div class="table-container has-mobile-cards">
            <table class={classes}>
                {#if head}<thead>{@render head()}</thead>{/if}
                {#if foot}<tfoot>{@render foot()}</tfoot>{/if}
                <tbody>{@render children?.()}</tbody>
            </table>
        </div>
    {:else}
        <div class="table-wrapper has-mobile-cards">
            <table class={classes}>
                {#if head}<thead>{@render head()}</thead>{/if}
                {#if foot}<tfoot>{@render foot()}</tfoot>{/if}
                <tbody>{@render children?.()}</tbody>
            </table>
        </div>
    {/if}
    {#if pagination}
        <div class="table-pagination notification">
            {@render pagination()}
        </div>
    {/if}
</div>
