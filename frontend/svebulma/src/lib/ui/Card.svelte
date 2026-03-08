<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        title?: string;
        icon?: string;
        hasTable?: boolean;
        header?: Snippet;
        headerAction?: Snippet;
        footer?: Snippet;
        children?: Snippet;
    }

    let { title, icon, hasTable = false, header, headerAction, footer, children }: Props = $props();

    const cardClasses = $derived([
        'card',
        hasTable ? 'has-table' : '',
    ].filter(Boolean).join(' '));
</script>

<div class={cardClasses}>
    {#if header || title}
        <header class="card-header">
            {#if header}
                {@render header()}
            {:else}
                <p class="card-header-title">
                    {#if icon}
                        <span class="icon"><i class="mdi mdi-{icon}"></i></span>
                    {/if}
                    {title}
                </p>
            {/if}
            {#if headerAction}
                <span class="card-header-icon">
                    {@render headerAction()}
                </span>
            {/if}
        </header>
    {/if}

    <div class="card-content">
        {@render children?.()}
    </div>

    {#if footer}
        <div class="card-footer">
            {@render footer()}
        </div>
    {/if}
</div>
