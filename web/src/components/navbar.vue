<template>
  <div class="header">
    <router-link :to="{ name: 'playground' }">
      <img class="logo" :src="require('@/assets/logo.png')"/>
    </router-link>
    <div class="right-menu">
      <el-menu
          default-active="1"
          class="menu"
          mode="horizontal"
          @select="handleSelect">
         <el-submenu index="1" v-if="isWalletConnected">
            <template slot="title">
                <span class="user">{{ truncatedWalletAddress }}</span>
            </template>
           <el-menu-item index="profile">
              <font-awesome-icon size="lg" icon="user-alt" style="margin-left:30px"></font-awesome-icon>
              <span style="margin-left:20px">{{$t('dashboard.navbar.personal_profile')}}</span>
           </el-menu-item>
           <div class="divider"></div>
           <el-menu-item index="logout">
              <font-awesome-icon size="lg" icon="sign-out-alt" style="margin-left:30px"></font-awesome-icon>
              <span style="margin-left:20px">{{$t('dashboard.navbar.logout')}}</span>
           </el-menu-item>
         </el-submenu>
         <el-menu-item index="1" v-else @click="showWalletModal = true">
            <span>{{$t('dashboard.navbar.connect_wallet')}}</span>
         </el-menu-item>
      </el-menu>
    </div>

    <!-- Wallet Connection Modal -->
 <el-dialog
    title="Connect Wallet"
    :visible.sync="showWalletModal"
    width="30%"
    custom-class="dark-theme-dialog">
    <div class="wallet-options">
      <el-button @click="connectWallet('metamask')" class="wallet-button">
        <img src="https://upload.wikimedia.org/wikipedia/commons/3/36/MetaMask_Fox.svg" alt="MetaMask" class="wallet-icon" />
        MetaMask
      </el-button>
      <el-button @click="connectWallet('walletconnect')" class="wallet-button">
        <img src="https://upload.wikimedia.org/wikipedia/commons/1/13/Walletconnect-logo.png" alt="WalletConnect" class="wallet-icon" />
        WalletConnect
      </el-button>
    </div>
  </el-dialog>
  </div>
</template>

<script>
import UserAPI from '@/api/v1/users'
import { mapState } from 'vuex'
import ErrorMessage from '@/model/errorMessage'
import { Message } from "element-ui"
import freezable from '@/mixins/freezable'

export default {
  name: "Navbar",
  mixins: [freezable],
  data() {
    return {
      walletAddress: '',
      showWalletModal: false,
    };
  },
  computed: {
    ...mapState({
      user: state => state.user
    }),
    isWalletConnected() {
      return this.walletAddress !== ''
    },
    truncatedWalletAddress() {
      if (this.walletAddress) {
        return `${this.walletAddress.slice(0, 6)}...${this.walletAddress.slice(-4)}`
      }
      return ''
    }
  },
  methods: {
    async connectWallet(walletType) {
      // Implement wallet connection logic here
      // This is a placeholder implementation
      try {
        let address;
        if (walletType === 'metamask') {
          // Connect to MetaMask
          if (typeof window.ethereum !== 'undefined') {
            const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
            address = accounts[0];
          } else {
            throw new Error('MetaMask is not installed');
          }
        } else if (walletType === 'walletconnect') {
          // Implement WalletConnect logic
          // You might need to install and import WalletConnect library
          throw new Error('WalletConnect not implemented');
        }

        if (address) {
          this.walletAddress = address;
          this.showWalletModal = false;
          this.login();
        }
      } catch (error) {
        console.error('Failed to connect wallet:', error);
        Message({ message: 'Failed to connect wallet: ' + error.message, type: 'error', duration: 3 * 1000 });
      }
    },
    login() {
      UserAPI.login(this.walletAddress, 'wallet').then((response) => {
        if(response.approve_status == this.$appGlobal.constants.USER_APPROVE_STATUS_REGISTED) {
          this.$router.push({ name: "post-regist" });
        } else {
          this.$store.commit("user/SET_USER", response);
        }
      }).catch((error) => {
        this.$errorMessage(error, (error) => {
          let errorMessage = error.response.data.message
          let showMessage = ErrorMessage.getLocalMessage(errorMessage)
          if (showMessage) {
            Message({ message: showMessage, type: 'error', duration: 3 * 1000})
          }
        })
      })
    },
    handleSelect(key, keyPath) {
       switch (key) {
          case 'logout':
             this.logout()
             break;
          case 'profile':
             this.$router.push({name:'profile'})
             break;
       }
    },
    logout() {
      // Implement logout logic
      this.walletAddress = '';
      this.$store.commit("user/CLEAR_USER");
      this.$router.push({ name: "login" });
    },
  },
};
</script>

<style lang="stylus" scoped>
.divider {
   margin 10px 30px
   height 1px
   background black
   opacity 0.1
}
.menu {
   background transparent
   border none
  .el-menu-item {
     span {
       color white
     }
     border-bottom 2px solid #151b23

     &:hover {
      span {
       color black
      }
       border-bottom 2px solid purple
     }
   }
}
.header {
  height 60px
  font-size 22px
  display flex
  justify-content space-between
  align-items center
  flex-shrink 0
  background-color #151b23
  border-bottom 1px solid #EBEEF5
  .logo {
    height 28px
    font-size 20px
    line-height 50px
    margin-left 30px
    color black
  }

  .right-menu {
    padding-right 93px
    .el-dropdown {
      cursor pointer
      background-color #151b23
    }
  }

  a,
  .user {
    font-size 16px
    margin 0px 10px
    text-decoration none
  }

  .menu-link {
    margin-right 30px
  }
}

.wallet-options {
  display flex
  flex-direction column
  align-items center
  
  .el-button {
    margin 10px 0
    width 200px
  }
}

::v-deep .el-dialog {
  background-color: #1a1a1a !important;
  border: 2px solid #4a0e4e

  ::v-deep .el-dialog__header {
    background-color: #2a2a2a
    padding: 20px
    
    ::v-deep .el-dialog__title {
      color: #ffffff
      font-size: 18px
      font-weight: bold
    }
  }

  ::v-deep .el-dialog__body {
    background-color: #1a1a1a
    padding: 30px
  }

  ::v-deep .el-dialog__headerbtn .el-dialog__close {
    color: #ffffff
  }
}

.wallet-options {
  display: flex
  flex-direction: column
  align-items: center
  
  .wallet-button {
    margin: 10px 0
    width: 100%
    height: 50px
    background-color: #2a2a2a
    border: 1px solid #4a0e4e
    color: #ffffff
    font-size: 16px
    display: flex
    align-items: center
    justify-content: flex-start
    padding: 0 20px
    transition: all 0.3s ease

    &:hover {
      background-color: #3a3a3a
      border-color: #6a2a6e
    }

    .wallet-icon {
      width: 24px
      height: 24px
      margin-right: 10px
    }
  }
}
</style>