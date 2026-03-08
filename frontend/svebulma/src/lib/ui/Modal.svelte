<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        open?: boolean;
        title?: string;
        header?: Snippet;
        footer?: Snippet;
        children?: Snippet;
        onclose?: () => void;
    }

    let { open = false, title, header, footer, children, onclose }: Props = $props();

    function close() {
        onclose?.();
    }
</script>

<div class={['modal', open ? 'is-active' : ''].filter(Boolean).join(' ')}>
    <div class="modal-background" onclick={close} role="presentation"></div>
    <div class="modal-card">
        <header class="modal-card-head">
            {#if header}
                {@render header()}
            {:else}
                <p class="modal-card-title">{title ?? ''}</p>
            {/if}
            <button class="delete" aria-label="close" onclick={close}></button>
        </header>

        <section class="modal-card-body">
            {@render children?.()}
        </section>

        {#if footer}
            <footer class="modal-card-foot">
                {@render footer()}
            </footer>
        {/if}
    </div>
</div>
