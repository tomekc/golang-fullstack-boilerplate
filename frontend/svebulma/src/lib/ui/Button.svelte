<script lang="ts">
    import type { Snippet } from 'svelte';

    type Color = 'primary' | 'link' | 'info' | 'success' | 'warning' | 'danger' | 'dark' | 'light' | 'white' | 'ghost';
    type Size = 'small' | 'normal' | 'medium' | 'large';

    interface Props {
        color?: Color;
        size?: Size;
        outlined?: boolean;
        inverted?: boolean;
        rounded?: boolean;
        loading?: boolean;
        fullwidth?: boolean;
        disabled?: boolean;
        type?: 'button' | 'submit' | 'reset';
        href?: string;
        onclick?: (e: MouseEvent) => unknown;
        children?: Snippet;
        [key: string]: unknown;
    }

    let {
        color,
        size,
        outlined = false,
        inverted = false,
        rounded = false,
        loading = false,
        fullwidth = false,
        disabled = false,
        type = 'button',
        href,
        onclick,
        children,
        ...rest
    }: Props = $props();

    const classes = $derived([
        'button',
        color ? `is-${color}` : '',
        size && size !== 'normal' ? `is-${size}` : '',
        outlined ? 'is-outlined' : '',
        inverted ? 'is-inverted' : '',
        rounded ? 'is-rounded' : '',
        loading ? 'is-loading' : '',
        fullwidth ? 'is-fullwidth' : '',
    ].filter(Boolean).join(' '));
</script>

{#if href}
    <a {href} class={classes} {...rest}>
        {@render children?.()}
    </a>
{:else}
    <button class={classes} {type} {disabled} {onclick} {...rest}>
        {@render children?.()}
    </button>
{/if}
