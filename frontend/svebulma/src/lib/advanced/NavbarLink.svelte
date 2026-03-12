<script lang="ts">
    interface Props {
        icon: string;
        label: string;
        href?: string;
        title?: string;
        hasDivider?: boolean;
        desktopIconOnly?: boolean;
    }

    let {
        icon,
        label,
        href = undefined,
        title = undefined,
        hasDivider = false,
        desktopIconOnly = false,
    }: Props = $props();
</script>

<a
    {title}
    class="navbar-item"
    class:has-divider={hasDivider}
    class:is-desktop-icon-only={desktopIconOnly}
    href={href}
>
    <span class="icon"><i class="mdi mdi-{icon}"></i></span>
    <span>{label}</span>
</a>
