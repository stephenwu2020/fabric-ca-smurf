<template>
  <div class="home">
    <div>
      <h1>CA Info</h1>
      <p>
        <span>CAChain</span>
        <span>
          {{ this.cainfo.CAChain }}
        </span>
      </p>
      <p>
        <span>CAName</span>
        <span>
          {{ this.cainfo.CAName }}
        </span>
      </p>
      <p>
        <span>IssuerPublicKeyC</span>
        <span>
          {{ this.cainfo.IssuerPublicKey }}
        </span>
      </p>
      <p>
        <span>IssuerRevocationPublicKey</span>
        <span>
          {{ this.cainfo.IssuerRevocationPublicKey }}
        </span>
      </p>
      <p>
        <span>Version</span>
        <span>
          {{ this.cainfo.Version }}
        </span>
      </p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      cainfo: {},
    };
  },
  name: "home",
  created() {
    this.$axios
      .get("http://127.0.0.1:8080/api/v1/cainfo")
      .then((res) => {
        this.cainfo = res.result;
      })
      .catch((err) => {
        console.error(err);
      });
  },
};
</script>

<style lang="postcss" scoped>
.home {
  h1 {
    text-align: left;
  }
  p {
    display: flex;
    margin: 0;
    span {
      overflow-x: scroll;
      display: inline-block;
      margin-right: 100px;
      text-align: left;
      &:nth-child(1) {
        width: 300px;
      }
      &:nth-child(2) {
        width: 600px;
      }
    }
  }
}
</style>
