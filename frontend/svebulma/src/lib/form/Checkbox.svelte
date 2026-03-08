<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        checked?: boolean;
        disabled?: boolean;
        required?: boolean;
        name?: string;
        id?: string;
        value?: string;
        onchange?: (e: Event) => void;
        children?: Snippet;
    }

    let {
        checked = $bindable(false),
        disabled = false,
        required = false,
        name,
        id,
        value,
        onchange,
        children,
    }: Props = $props();
</script>

<label class={['checkbox', disabled ? '' : ''].filter(Boolean).join(' ')}>
    <input
        type="checkbox"
        bind:checked
        {disabled}
        {required}
        {name}
        {id}
        {value}
        {onchange}
    />
    {#if children}
        &nbsp;{@render children()}
    {/if}
</label>
