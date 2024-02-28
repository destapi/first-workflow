# Practising with github actions
## Workflows, jobs and steps

_Workflows_ belong to a _Repository_ (1 or more)
> Triggered upon _Events_

Workflow triggers
> push - pushing a commit  
> pull_request - pull request action  
> create - createing a branch or tag  
> fork - forking a repository  
> issues - opening or deleting an issue  
> issue_comment - issue or pull request comment  
> watch - repository was starred  
> discussion - discussion was created or deleted  
> etc  
> workflow_dispatch - manually triggered workflow  
> repository_dispatch - REST API request triggered action  
> schedule - workflow is scheduled  
> workflow_call - can be called by other workflows  

_Jobs_ belong to a _Workflow_ (1 or more)
> Executed in a _Runner_ (execution environment)  
> Every _Job_ has its own _Runner_, and the _Runners_ can be different for each _Job_
> _Jobs_ can run either in parallel (default) or sequantially  
> Can be run conditionally  

_Steps_ belong to a _Job_ (1 or more)
> Execute an _Action_ (shell script)  
> _Steps_ are executed in sequential order (not in parallel)  
> Can be executed conditionally as well  

