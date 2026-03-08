<script lang="ts">
    type Color = 'primary' | 'link' | 'info' | 'success' | 'warning' | 'danger';
    type Size = 'small' | 'normal' | 'medium' | 'large';

    interface Props {
        value?: string;
        placeholder?: string;
        color?: Color;
        size?: Size;
        rows?: number;
        fixedSize?: boolean;
        loading?: boolean;
        disabled?: boolean;
        readonly?: boolean;
        required?: boolean;
        name?: string;
        id?: string;
        onchange?: (e: Event) => void;
        oninput?: (e: Event) => void;
    }

    let {
        value = $bindable(''),
        placeholder,
        color,
        size,
        rows = 4,
        fixedSize = false,
        loading = false,
        disabled = false,
        readonly = false,
        required = false,
        name,
        id,
        onchange,
        oninput,
    }: Props = $props();

    const classes = $derived([
        'textarea',
        color ? `is-${color}` : '',
        size && size !== 'normal' ? `is-${size}` : '',
        loading ? 'is-loading' : '',
        fixedSize ? 'has-fixed-size' : '',
    ].filter(Boolean).join(' '));
</script>

<textarea
    class={classes}
    bind:value
    {placeholder}
    {rows}
    {disabled}
    {readonly}
    {required}
    {name}
    {id}
    {onchange}
    {oninput}
></textarea>
