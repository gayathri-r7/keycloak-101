module.exports = () => {
    const Service = require('@vue/cli-service/lib/Service')
    const service = new Service(process.env.VUE_CLI_CONTEXT || process.cwd())
  
    const { error } = require('@vue/cli-shared-utils')
    service.run('serve').catch(err => {
      error(err)
      process.exit(1)
    })
  }
  