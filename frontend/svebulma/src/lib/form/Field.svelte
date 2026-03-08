<script lang="ts">
    import type { Snippet } from 'svelte';

    type State = 'danger' | 'success' | 'warning' | 'info';

    interface Props {
        label?: string;
        for?: string;
        help?: string;
        state?: State;
        horizontal?: boolean;
        grouped?: boolean;
        children?: Snippet;
    }

    let { label, for: labelFor, help, state, horizontal = false, grouped = false, children }: Props = $props();
</script>

<div class={['field', horizontal ? 'is-horizontal' : '', grouped ? 'is-grouped' : ''].filter(Boolean).join(' ')}>
    {#if horizontal}
        {#if label}
            <div class="field-label is-normal">
                <label class="label" for={labelFor}>{label}</label>
            </div>
        {/if}
        <div class="field-body">
            <div class="field">
                <div class="control">
                    {@render children?.()}
                </div>
                {#if help}
                    <p class={['help', state ? `is-${state}` : ''].filter(Boolean).join(' ')}>{help}</p>
                {/if}
            </div>
        </div>
    {:else}
        {#if label}
            <label class="label" for={labelFor}>{label}</label>
        {/if}
        <div class="control">
            {@render children?.()}
        </div>
        {#if help}
            <p class={['help', state ? `is-${state}` : ''].filter(Boolean).join(' ')}>{help}</p>
        {/if}
    {/if}
</div>
