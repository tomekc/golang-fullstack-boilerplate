<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        label: string;
        children: Snippet;
    }

    let { label, children }: Props = $props();
</script>

<p class="menu-label">{label}</p>
<ul class="menu-list">
    {@render children()}
</ul>
