<template>
<!--  "class="ivu-table"-->
  <div class="myTable">
    <Table :border="showBorder" :stripe="showStripe" :show-header="showHeader" :height="fixedHeader ? 250 : ''"
           :size="tableSize" :loading="loading" :data="tableData" :columns="tableColumns" :max-height="700"></Table>



  </div>
</template>

<script>
export default {
    name: 'myTable',
    data () {
        return {
            showBorder: true,
            showStripe: false,
            showHeader: true,
            showIndex: false,
            showCheckbox: false,
            fixedHeader: false,
            tableSize: 'default'
        }
    },
    props:{
      tableData : Array,
      loading: Boolean,
    },
    computed: {
        tableColumns () {
            let columns = [];
            if (this.showCheckbox) {
                columns.push({
                    type: 'selection',
                    width: 60,
                    align: 'center'
                })
            }
            if (this.showIndex) {
                columns.push({
                    type: 'index',
                    width: 60,
                    align: 'center'
                })
            }
            columns.push({
                title: '地标IP',
                key: 'Ip',
                sortable: true
            });
            columns.push({
                title: '跳数',
                key: 'JumpNum',
                sortable: true,
            });
            columns.push({
                title: '地标类型',
                key: 'Typename',
                sortable: true,
            });
            columns.push({
                title: '地标名称',
                key: 'Name',
            });
            columns.push({
                title: '国家',
                key: 'Country',
            });
            columns.push({
                title: '省份',
                key: 'Province',
            });
            columns.push({
                title: '城市',
                key: 'City',
            });
            return columns;
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  /*.myTable{*/
  /*  margin: 30px;*/
  /*}*/
</style>
