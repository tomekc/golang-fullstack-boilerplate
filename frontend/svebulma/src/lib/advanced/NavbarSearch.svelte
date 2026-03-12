<script lang="ts">
    interface Props {
        placeholder?: string;
    }

    let { placeholder = 'Search everywhere...' }: Props = $props();
</script>

<div class="navbar-item has-control">
    <div class="control"><input {placeholder} class="input"></div>
</div>
