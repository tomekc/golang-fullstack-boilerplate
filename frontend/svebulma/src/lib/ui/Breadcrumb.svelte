<script lang="ts">
    type Alignment = 'centered' | 'right';
    type Separator = 'arrow' | 'bullet' | 'dot' | 'succeeds';
    type Size = 'small' | 'medium' | 'large';

    interface BreadcrumbItem {
        label: string;
        href?: string;
        active?: boolean;
    }

    interface Props {
        items: BreadcrumbItem[];
        alignment?: Alignment;
        separator?: Separator;
        size?: Size;
    }

    let { items, alignment, separator, size }: Props = $props();

    const classes = $derived([
        'breadcrumb',
        alignment ? `is-${alignment}` : '',
        separator ? `has-${separator}-separator` : '',
        size ? `is-${size}` : '',
    ].filter(Boolean).join(' '));
</script>

<nav class={classes} aria-label="breadcrumbs">
    <ul>
        {#each items as item}
            <li class:is-active={item.active}>
                {#if item.href && !item.active}
                    <a href={item.href}>{item.label}</a>
                {:else}
                    <a href={item.href ?? '#'} aria-current={item.active ? 'page' : undefined}>{item.label}</a>
                {/if}
            </li>
        {/each}
    </ul>
</nav>
