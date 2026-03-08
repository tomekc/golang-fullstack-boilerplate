<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        title: string;
        actions?: Snippet;
    }

    let { title, actions }: Props = $props();
</script>

<section class="hero is-hero-bar">
    <div class="hero-body">
        <div class="level">
            <div class="level-left">
                <div class="level-item">
                    <h1 class="title">{title}</h1>
                </div>
            </div>
            {#if actions}
                <div class="level-right">
                    <div class="level-item">
                        {@render actions()}
                    </div>
                </div>
            {/if}
        </div>
    </div>
</section>
