<template>
   <div>
      <nav class="navbar navbar-expand-lg navbar-light bg-light title-app">
         <div style="padding-left: 5%" class="collapse navbar-collapse" id="navbarTogglerDemo02">
            <h2 class="text-left">Todo Application</h2>
         </div>
         <div style="padding-right: 5%"><b-button variant="outline-primary">Logout</b-button></div>
      </nav>
      <div id="app" class="container mt-5">
         <b-card class="mt-3">
            <b-row>
               <b-col md="10">
                  <b-input v-model="model.task" placeholder="Add new task..."></b-input>
                  <b-form-textarea class="mt-3"
                     id="textarea"
                     v-model="model.description"
                     placeholder="Enter description..."
                     rows="3"
                     max-rows="6">
                  </b-form-textarea>
               </b-col>
               <b-col md="2">
                  <b-button variant="outline-primary" @click="createTodo()">Add Task</b-button>
               </b-col>
            </b-row>
         </b-card>
         <b-row class="mt-5">
            <b-col>
               <h2>Task list</h2>
               <b-table v-if="resp && resp.length != 0" ref="table" stacked="md" :items="resp" :fields="fields" class="table-style" outlined>
                  <template v-slot:cell(task)="resp">
                     {{resp.item.task}}
                  </template>
                  <template v-slot:cell(status)="resp">
                     <code>{{resp.item.status}}</code>
                  </template>
                  <template v-slot:cell(actions)="resp">
                     <b-button variant="outline-primary" v-b-modal="'view-todo-' + resp.item.id">View</b-button>
                     &nbsp;&nbsp;
                     <b-button style="padding-right: 10px" class="ml-3" variant="outline-primary" v-b-modal="'edit-todo-' + resp.item.id">Edit</b-button>
                     &nbsp;&nbsp;
                     <b-button variant="outline-primary" @click="deleteTodo(resp.item.id)">Delete</b-button>

                     <b-modal :id="'edit-todo-' + resp.item.id" title="Edit Todo" okTitle="Save" size="lg" @ok="saveTodo(resp.item.id)">
                        <TodoForm :todoID="resp.item.id" @saveTodoData="todoSavePayload($event)" formType="edit"></TodoForm>
                    </b-modal>
                    <b-modal :id="'view-todo-' + resp.item.id" title="View Todo" size="lg">
                        <TodoForm :todo="todo" :todoID="resp.item.id" formType="view"></TodoForm>
                    </b-modal>

                  </template>
               </b-table>
                <div v-else-if="error && error != ''" class="alert alert-danger alert-dismissible fade show text-center" role="alert">
                  <strong>Error: </strong> {{error}}
                </div>
                <div v-else class="alert alert-info alert-dismissible fade show text-center" role="alert">
                  <strong>Info: </strong> No tasks found.
                </div>
            </b-col>
         </b-row>
      </div>
   </div>
</template>

<script>
 import axios from "axios";
  export default {
    name: 'TodoApp',
    data() {
      return {
        entry: {},
        model: {},
        error: "",
        todo: {},
        resp: [],
        fields: [
          {
            key: 'task',
            label: 'Task'
          },
          {
            key: 'status',
            label: 'Status'
          },
          {
            key: 'actions',
            label: 'Actions'
          }, 
        ]
      }
    },
    components: {
      'TodoForm': () => import('@/components/Forms/TodoForm.vue'),
    },
    created() {
        this.listTodos();
    },
    methods: {
      makeToast(variant = null, message)
      {
        this.$bvToast.toast(`${message}`,
        {
          title: `Variant ${variant || 'default'}`,
          variant: variant,
          solid: true,
          autoHideDelay: 5000,
        });
      },
      todoSavePayload(data){
        this.todo = data;
      },
      // Create Todo
      createTodo()
      {
        var url = "/todo";
        var model = JSON.parse(JSON.stringify(this.model));
        return axios(
        {
          url: url,
          method: 'POST',
          data: model
        }).then(result =>
        {
          this.listTodos();
          this.makeToast('success', "Task added successfully.");
          return result;
        }).catch(error =>
        {
          console.log(error);
          throw error;
        });
      },
      saveTodo(id)
      {
        var entry = {}
        Object.assign(entry, this.todo);
        var url = "/todo/" + id;
        var model = JSON.parse(JSON.stringify(entry));
        return axios(
        {
          url: url,
          method: 'PUT',
          data: model
        }).then(result =>
        {
          this.listTodos();
          this.makeToast('success', "Task saved successfully.");
          return result;
        }).catch(error =>
        {
          console.log(error);
          throw error;
        });
      },
      // List Todos
      listTodos()
      {
        var url = "/todos";
        return axios(
        {
          url: url,
          method: 'GET',
        }).then(result => {
          console.log(typeof(result.data))
          this.resp = result.data;
          return result;
        }).catch(error => {
          console.log(error);
          this.error = error;
          throw error;
        });
      },
      deleteTodo(id)
      {
        var url = "/todo/" + id;
        return axios(
        {
          url: url,
          method: 'DELETE',
        }).then(result => {
          this.makeToast('success', "Task removed successfully.");
          this.listTodos();
          return result;
        }).catch(error =>
        {
          console.log(error);
          this.error = error;
          throw error;
        });
      },
    },
  }
</script>
<style scoped>
.title-app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 5px;
}
</style>