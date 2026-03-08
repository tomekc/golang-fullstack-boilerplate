<script lang="ts">
    import type { Snippet } from 'svelte';

    type Color = 'primary' | 'link' | 'info' | 'success' | 'warning' | 'danger';
    type Size = 'small' | 'normal' | 'medium' | 'large';

    interface SelectOption {
        value: string | number;
        label: string;
        disabled?: boolean;
    }

    interface Props {
        value?: string | number;
        options?: SelectOption[];
        placeholder?: string;
        color?: Color;
        size?: Size;
        rounded?: boolean;
        loading?: boolean;
        multiple?: boolean;
        disabled?: boolean;
        required?: boolean;
        name?: string;
        id?: string;
        children?: Snippet;
        onchange?: (e: Event) => void;
    }

    let {
        value = $bindable(''),
        options,
        placeholder,
        color,
        size,
        rounded = false,
        loading = false,
        multiple = false,
        disabled = false,
        required = false,
        name,
        id,
        children,
        onchange,
    }: Props = $props();

    const wrapperClasses = $derived([
        'select',
        color ? `is-${color}` : '',
        size && size !== 'normal' ? `is-${size}` : '',
        rounded ? 'is-rounded' : '',
        loading ? 'is-loading' : '',
        multiple ? 'is-multiple' : '',
    ].filter(Boolean).join(' '));
</script>

<div class={wrapperClasses}>
    {#if multiple}
        <select multiple bind:value {disabled} {required} {name} {id} {onchange}>
            {#if options}
                {#each options as opt}
                    <option value={opt.value} disabled={opt.disabled}>{opt.label}</option>
                {/each}
            {:else}
                {@render children?.()}
            {/if}
        </select>
    {:else}
        <select bind:value {disabled} {required} {name} {id} {onchange}>
            {#if placeholder}
                <option value="" disabled selected={!value}>{placeholder}</option>
            {/if}
            {#if options}
                {#each options as opt}
                    <option value={opt.value} disabled={opt.disabled}>{opt.label}</option>
                {/each}
            {:else}
                {@render children?.()}
            {/if}
        </select>
    {/if}
</div>
