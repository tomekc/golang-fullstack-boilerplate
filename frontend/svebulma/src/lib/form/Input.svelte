<script lang="ts">
    type InputType = 'text' | 'email' | 'password' | 'number' | 'tel' | 'url' | 'search' | 'date' | 'time' | 'datetime-local';
    type Color = 'primary' | 'link' | 'info' | 'success' | 'warning' | 'danger';
    type Size = 'small' | 'normal' | 'medium' | 'large';

    interface Props {
        value?: string | number;
        type?: InputType;
        placeholder?: string;
        color?: Color;
        size?: Size;
        rounded?: boolean;
        loading?: boolean;
        disabled?: boolean;
        readonly?: boolean;
        required?: boolean;
        name?: string;
        id?: string;
        onchange?: (e: Event) => void;
        oninput?: (e: Event) => void;
        [key: string]: unknown;
    }

    let {
        value = $bindable(''),
        type = 'text',
        placeholder,
        color,
        size,
        rounded = false,
        loading = false,
        disabled = false,
        readonly = false,
        required = false,
        name,
        id,
        onchange,
        oninput,
        ...rest
    }: Props = $props();

    const classes = $derived([
        'input',
        color ? `is-${color}` : '',
        size && size !== 'normal' ? `is-${size}` : '',
        rounded ? 'is-rounded' : '',
        loading ? 'is-loading' : '',
    ].filter(Boolean).join(' '));
</script>

<input
    class={classes}
    {type}
    bind:value
    {placeholder}
    {disabled}
    {readonly}
    {required}
    {name}
    {id}
    {onchange}
    {oninput}
    {...rest}
/>
