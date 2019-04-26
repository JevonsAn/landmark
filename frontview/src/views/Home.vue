<template>
  <div class="layout">
    <Layout>
      <Header>
        <div class="layout-title">地 标 查 询</div>
      </Header>
      <Layout :style="{padding: '0 50px'}">
        <Breadcrumb :style="{margin: '16px 0'}">
        </Breadcrumb>
        <Content :style="{padding: '24px 0', minHeight: '700px', background: '#fff'}">
          <Layout>
            <Content :style="{padding: '24px', minHeight: '700px', background: '#fff'}">

              <div class="home">
                <h1>最少跳数地标点查询</h1>
                <Divider />
                <div class="form">
                  <Row class="form-line">
                    <Col class="form-text" span="2"> TopK：</Col>
                    <Col span="11">
                      <Input style="width: 20%" v-model="topNum" placeholder="请输入整数" size="large"/>
                    </Col>
                  </Row>
                  <br/>
                  <Row class="form-line">
                    <Col class="form-text" span="2"> IP地址：</Col>
                    <Col span="11">
                      <Input style="width: 50%" v-model="address" placeholder="请输入IP地址" size="large" search enter-button @on-search="mysearch"/>
                    </Col>
                  </Row>
                </div>
                <Divider />
                <div class="table">
                  <myTable :tableData="tableData" :loading="loading"/>
                </div>
              </div>

            </Content>
          </Layout>
        </Content>
      </Layout>
      <Footer class="layout-footer-center">2019 &copy; HIT-NIS</Footer>
    </Layout>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios';
import myTable from '@/components/Table.vue'

export default {
  name: 'home',
  data: function () {
    return {
      address: "",
      tableData: [],
      loading: false,
      topNum: 3
    }
  },
  methods:{
    mysearch: function () {
        console.log(this.address);
        this.loading = true;
        let vm = this;
        axios.get('http://10.10.11.132:8888/nearLM?n=' + this.topNum + '&ip=' + this.address)
          .then(function (response) {
            vm.tableData = response.data;
            console.log(response.data);
            vm.loading = false;
          })
          .catch(function (error) {
            console.log('Error! Could not reach the API. ' + error)
          })
    }
  },
  components: {
    myTable
  }
}
</script>

<style>
  .home {
    text-align: left;
    margin: 0px 30px 30px;
  }
  /*.home .form-line{*/
  /*  padding-bottom: 18px;*/
  /*}*/
  .home .form-text{
    font-size: large;
    padding-left: 10px;
    padding-top: 4px;
  }
  .layout{
      border: 1px solid #d7dde4;
      background: #f5f7f9;
      position: relative;
      border-radius: 4px;
      overflow: hidden;
  }
  .layout-title{
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","楷体","微软雅黑",Arial,sans-serif;

    text-align: left;
      font-size: xx-large;
      color: #d7dde4;
      /*width: 100px;*/
      /*height: 30px;*/
      /*background: #5b6270;*/
      /*border-radius: 3px;*/
      /*float: left;*/
      /*position: relative;*/
      top: 15px;
      left: 20px;
  }
  .layout-footer-center{
      text-align: center;
  }
</style>