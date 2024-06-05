<script>
  /**
   * @type {{user: string, data: string}[]}
   */
  let messages = [];
  /**
   * @type {string}
   */
  let username =
    Math.floor(Math.random() * 10).toString() +
    Math.floor(Math.random() * 10).toString() +
    Math.floor(Math.random() * 10).toString() +
    Math.floor(Math.random() * 10).toString() +
    Math.floor(Math.random() * 10).toString() +
    Math.floor(Math.random() * 10).toString() +
    Math.floor(Math.random() * 10).toString();
  /**
   * @type {string}
   */
  let text;
  // let url = "http://localhost:8080/messages";
  let url = "https://e7ebb400bf917d.lhr.life/messages";

  let fetchChats = () => {
    fetch(url)
      .then((res) => res.json())
      .then((result) => {
        messages = result;
      });
  };
  let startFetchInterval = () => {
    setInterval(() => {
      fetchChats();
    }, 1000);
  };
  let postChat = () => {
    let newChat = { user: username, data: text };
    fetch(url, {
      method: "post",
      body: JSON.stringify(newChat),
    });
    messages.push(newChat);
  };
</script>

<main>
  <div class="flex">
    <label for="data"
      >내용<input name="data" type="text" bind:value={text} /></label
    >

    <button on:click={postChat}>확인</button>
  </div>
  <button on:click={startFetchInterval}>인터벌 시작</button>
  {#each messages as message}
    {#if username === message.user}
      <p class="text me">{message.data}</p>
    {:else}
      <p class="text other">({message.user}) {message.data}</p>
    {/if}
  {/each}
</main>

<style>
  .flex {
    margin-top: 10px;
    display: flex;
    flex-direction: column;
    width: fit-content;
  }
  .text {
    width: fit-content;
    border-radius: 10px;
    padding: 10px;
    color: white;
  }
  .other {
    background-color: #777;
  }
  .me {
    background-color: #333;
    margin-left: auto;
  }
</style>
