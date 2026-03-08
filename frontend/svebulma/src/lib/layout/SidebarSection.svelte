<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        title: string;
        children?: Snippet;
    }

    let { title, children }: Props = $props();
</script>

<p class="menu-label">{title}</p>
<ul class="menu-list">
    {@render children?.()}
</ul>
