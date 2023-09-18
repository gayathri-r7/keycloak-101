<template>
<div class="container">
    <b-row>
        <b-col md="4">
            <img src="@/assets/todo-icon.avif" alt="Italian Trulli" width="200" height="200">
        </b-col>
        <b-col md="8">
            <div class="input-group mb-1">
                <input :disabled="formType === 'view'" type="text" class="form-control" placeholder="Add new task..." aria-label="Add new task..." v-model="model.task" aria-describedby="basic-addon1">
            </div>
            <div class="input-group mb-1">
                <b-form-textarea :disabled="formType === 'view'" class="mt-3" id="textarea" v-model="model.description" placeholder="Enter description..." rows="3" max-rows="6"></b-form-textarea>
            </div>
            <div class="input-group mt-3 mb-3">
                <select :disabled="formType === 'view'" class="form-select" v-model="model.status" aria-label="Default select example">
                    <option value="In-Progress">In-Progress</option>
                    <option value="Completed">Completed</option>
                </select>
            </div>
        </b-col>
    </b-row>
</div>
</template>

<script>
import axios from "axios";
export default {
  name: 'TodoForm',
  props: ['todoID', 'formType'],
  data() {
      return {
        model: {}
      }
    },
    watch: {
        'model': function(val) {
            this.$emit('saveTodoData', val);
        },
    },
    created() {
        this.getTodo(this.todoID);
    },
    methods: {
      getTodo(id)
      {
        var url = "/todo/" + id;
        return axios({
          url: url,
          method: 'GET',
        }).then(result => {
          this.model = result.data;
          return result;
        }).catch(error => {
          console.log(error);
          this.error = error;
          throw error;
        });
      },
    },
}
</script>

<style scoped>
@media (min-width: 1025px) {
    .h-custom {
        height: 100vh !important;
    }
}

</style>