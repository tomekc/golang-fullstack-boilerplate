<script lang="ts">
    type Size = 'small' | 'normal' | 'medium' | 'large';
    type Color = 'primary' | 'link' | 'info' | 'success' | 'warning' | 'danger' | 'dark' | 'light' | 'white' | 'black' | 'grey' | 'grey-light' | 'grey-dark';

    interface Props {
        name: string;
        size?: Size;
        color?: Color;
        spin?: boolean;
        prefix?: string;
    }

    let { name, size, color, spin = false, prefix = 'fas' }: Props = $props();

    const spanClasses = $derived([
        'icon',
        size && size !== 'normal' ? `is-${size}` : '',
        color ? `has-text-${color}` : '',
    ].filter(Boolean).join(' '));

    const iconClasses = $derived([
        prefix,
        `fa-${name}`,
        spin ? 'fa-spin' : '',
    ].filter(Boolean).join(' '));
</script>

<span class={spanClasses}>
    <i class={iconClasses}></i>
</span>
