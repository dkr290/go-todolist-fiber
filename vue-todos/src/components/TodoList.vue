<!-- <template>


    <el-row>
    <el-col :span="12" :offset="7" style="width: 100%;">
        <h1>Todo List</h1>
        <todo-form @send-message="createTodo"></todo-form>
        <el-table :data="todos" style="width: 100%">
            <el-table-column prop="title" label="Title" width="320" />
            <el-table-column fixed="right" label="Operations" width="200" />
            <template #test="scope">
                <el-space wrap>
                    <el-switch
                        v-model="scope.row.completed"
                        @click="updateTodo(scope.row)"
                     />
                   
                </el-space>
            </template>
        </el-table>
    </el-col>
  </el-row>

</template> -->

<template>
    <el-row>
    <el-col :span="12" :offset="7" style="width: 100%;">
        <h1>Todo List</h1>
        <todo-form @send-message="createTodo"></todo-form>
    </el-col>
    </el-row>
    <el-table :data="todos" style="width: 100%">
        <el-table-column prop="title" label="Title" width="380" />
            <el-table-column fixed="right" label="Completed" width="100" />
            <el-table-column fixed="right" label="Delete" width="100" />
      <el-table-column fixed="right">
        <!-- <template #header>
          <el-input v-model="search" size="small" placeholder="Type to search" />
        </template> -->
        <template #default="scope">
            <el-space wrap>
                    <el-switch
                        v-model="scope.row.completed"
                        @click="updateTodo(scope.row)"
                     />
                   
                </el-space>
          <el-popconfirm width="30"
           title="Are you sure to delete this?"
           confirm-button-text="Yes"
           cancel0button-text="No"
           icon-color="red"
           @confirm="handleDelete(scope.row)">
        
             <template #reference>
             <el-button size="small"
            type="danger"
            
            
            >Delete</el-button>
            </template>
            </el-popconfirm>
          
        </template>
      </el-table-column>
    </el-table>
  </template>

<script lang="ts">
    import { ElMessage } from 'element-plus';
    import { Options, Vue } from 'vue-class-component';
    import { computed, ref } from 'vue'
    import TodoForm from './TodoForm.vue'

    interface Todo{
        id: number;
        title: string;
        completed: boolean;

    }


     
    @Options({

        components:{
           TodoForm,
        }
    })
     export default class TodoKList extends Vue{
        todos =[]; 
         async mounted(){
            await this.loadTodos();
         }

   
    
       async loadTodos(){
        const response = await this.axios.get(`http://localhost:8001/api/get_todos`)
        this.todos = response.data;
        
       }

        async createTodo(todo: any){
            console.log("Todo",todo)
            await this.axios.post(`http://localhost:8001/api/create_todos`,{
                title: todo.title,
                completed: todo.complated
            });
            ElMessage({
                message: "Todo Created",
                type: "success"
            })
            await this.loadTodos();
        }
        async updateTodo(todo: Todo) {
            console.log("Todo",todo)

            await this.axios.put(`http://localhost:8001/api/update_todos/${todo.id}`,{
                id: todo.id,
                title: todo.title,
                completed: todo.completed
            });
            ElMessage({
                message: "Todo Updated",
                type: "success"
            })
            await this.loadTodos();
        }

        async handleDelete(todo: Todo){

            await this.axios.delete(`http://localhost:8001/api/delete_todos/${todo.id}`);

            ElMessage({
                message: "Todo Deleted",
                type: "success"
            })
            await this.loadTodos();

        }

        cacnelDelete(){
            console.log("Cancel the delete")
        }

     }
</script>