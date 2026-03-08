<script lang="ts">
    type Color = 'primary' | 'link' | 'info' | 'success' | 'warning' | 'danger' | 'dark';
    type Size = 'small' | 'normal' | 'medium' | 'large';

    interface Props {
        value?: number;
        max?: number;
        color?: Color;
        size?: Size;
        indeterminate?: boolean;
    }

    let { value, max = 100, color, size, indeterminate = false }: Props = $props();

    const classes = $derived([
        'progress',
        color ? `is-${color}` : '',
        size && size !== 'normal' ? `is-${size}` : '',
    ].filter(Boolean).join(' '));
</script>

{#if indeterminate}
    <progress class={classes}></progress>
{:else}
    <progress class={classes} {value} {max}>{value}%</progress>
{/if}
