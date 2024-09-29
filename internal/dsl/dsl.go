package dsl

const DSLVersion = "1.0.0-alpha1"
const DSLSpec = `
$id: https://serverlessworkflow.io/schemas/1.0.0/workflow.yaml
$schema: https://json-schema.org/draft/2020-12/schema
description: Serverless Workflow DSL - Workflow Schema.
type: object
required: [ document, do ]
properties:
  document:
    type: object
    title: Document
    description: Documents the workflow.
    unevaluatedProperties: false
    properties:
      dsl:
        type: string
        pattern: ^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$
        title: WorkflowDSL
        description: The version of the DSL used by the workflow.
      namespace:
        type: string
        pattern: ^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$
        title: WorkflowNamespace
        description: The workflow's namespace.
      name:
        type: string
        pattern: ^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$
        title: WorkflowName
        description: The workflow's name.
      version:
        type: string
        pattern: ^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$  
        title: WorkflowVersion
        description: The workflow's semantic version.
      title:
        type: string
        title: WorkflowTitle
        description: The workflow's title.
      summary:
        type: string
        title: WorkflowSummary
        description: The workflow's Markdown summary.
      tags:
        type: object
        title: WorkflowTags
        description: A key/value mapping of the workflow's tags, if any.
        additionalProperties: true
      metadata:
        type: object
        title: WorkflowMetadata
        description: Holds additional information about the workflow.
        additionalProperties: true
    required: [ dsl, namespace, name, version ]
  input:
    $ref: '#/$defs/input'
    title: Input
    description: Configures the workflow's input.
  use:
    type: object
    title: Use
    description: Defines the workflow's reusable components.
    unevaluatedProperties: false
    properties:
      authentications:
        type: object
        title: UseAuthentications
        description: The workflow's reusable authentication policies.
        additionalProperties:
          $ref: '#/$defs/authenticationPolicy'
      errors:
        type: object
        title: UseErrors
        description: The workflow's reusable errors.
        additionalProperties:
          $ref: '#/$defs/error'
      extensions:
        type: array
        title: UseExtensions
        description: The workflow's extensions.
        items:
          type: object
          title: ExtensionItem
          minProperties: 1
          maxProperties: 1
          additionalProperties:
            $ref: '#/$defs/extension'
      functions:
        type: object
        title: UseFunctions
        description: The workflow's reusable functions.
        additionalProperties:
          $ref: '#/$defs/task'
      retries:
        type: object
        title: UseRetries
        description: The workflow's reusable retry policies.
        additionalProperties:
          $ref: '#/$defs/retryPolicy'
      secrets:
        type: array
        title: UseSecrets
        description: The workflow's reusable secrets.
        items:
          type: string
          description: The workflow's secrets.
      timeouts:
        type: object
        title: UseTimeouts
        description: The workflow's reusable timeouts.
        additionalProperties:
          $ref: '#/$defs/timeout'
  do:
    $ref: '#/$defs/taskList'
    title: Do
    description: Defines the task(s) the workflow must perform.
  timeout:
    oneOf:
      - $ref: '#/$defs/timeout'
        title: TimeoutDefinition
        description: The workflow's timeout configuration, if any.
      - type: string
        title: TimeoutReference
        description: The name of the workflow's timeout, if any.
  output:
    $ref: '#/$defs/output'
    title: Output
    description: Configures the workflow's output.
  schedule:
    type: object
    title: Schedule
    description: Schedules the workflow.
    unevaluatedProperties: false
    properties:
      every:
        $ref: '#/$defs/duration'
        title: ScheduleEvery
        description: Specifies the duration of the interval at which the workflow should be executed.
      cron:
        type: string
        title: ScheduleCron
        description: Specifies the schedule using a cron expression, e.g., '0 0 * * *' for daily at midnight.
      after:
        $ref: '#/$defs/duration'
        title: ScheduleAfter
        description: Specifies a delay duration that the workflow must wait before starting again after it completes.
      on:
        $ref: '#/$defs/eventConsumptionStrategy'
        title: ScheduleOn
        description: Specifies the events that trigger the workflow execution.
$defs:
  taskList:
    title: TaskList
    description: List of named tasks to perform.
    type: array
    items:
      type: object
      title: TaskItem
      minProperties: 1
      maxProperties: 1
      additionalProperties:
        $ref: '#/$defs/task'
  taskBase:
    type: object
    title: TaskBase
    description: An object inherited by all tasks.
    properties:
      if:
        type: string
        title: TaskBaseIf
        description: A runtime expression, if any, used to determine whether or not the task should be run.
      persistence:
        $ref: '#/$defs/persistence'
        title: TaskBasePersistence
        description: Configure the task's persistence.
      input:
        $ref: '#/$defs/input'
        title: TaskBaseInput
        description: Configure the task's input.
      output:
        $ref: '#/$defs/output'
        title: TaskBaseOutput
        description: Configure the task's output.
      export:
        $ref: '#/$defs/export'
        title: TaskBaseExport
        description: Export task output to context.
      timeout:
        oneOf:
          - $ref: '#/$defs/timeout'
            title: TaskTimeoutDefinition
            description: The task's timeout configuration, if any.
          - type: string
            title: TaskTimeoutReference
            description: The name of the task's timeout, if any.
      then:
        $ref: '#/$defs/flowDirective'
        title: TaskBaseThen
        description: The flow directive to be performed upon completion of the task.
      metadata:
        type: object
        title: TaskMetadata
        description: Holds additional information about the task.
        additionalProperties: true
  task:
    title: Task
    description: A discrete unit of work that contributes to achieving the overall objectives defined by the workflow.
    unevaluatedProperties: false
    oneOf:
      - $ref: '#/$defs/callTask'
      - $ref: '#/$defs/doTask'
      - $ref: '#/$defs/forkTask'
      - $ref: '#/$defs/emitTask'
      - $ref: '#/$defs/forTask'
      - $ref: '#/$defs/listenTask'
      - $ref: '#/$defs/raiseTask'
      - $ref: '#/$defs/runTask'
      - $ref: '#/$defs/setTask'
      - $ref: '#/$defs/switchTask'
      - $ref: '#/$defs/tryTask'
      - $ref: '#/$defs/waitTask'
  callTask:
    title: CallTask
    description: Defines the call to perform.
    oneOf:
      - title: CallAsyncAPI
        description: Defines the AsyncAPI call to perform.
        $ref: '#/$defs/taskBase'
        type: object
        required: [ call, with ]
        unevaluatedProperties: false
        properties:
          call:
            type: string
            const: asyncapi
          with:
            type: object
            title: AsyncApiArguments
            description: The Async API call arguments.
            properties:
              document:
                $ref: '#/$defs/externalResource'
                title: WithAsyncAPIDocument
                description: The document that defines the AsyncAPI operation to call.
              operationRef:
                type: string
                title: WithAsyncAPIOperation
                description: A reference to the AsyncAPI operation to call.
              server:
                type: string
                title: WithAsyncAPIServer
                description: A a reference to the server to call the specified AsyncAPI operation on. If not set, default to the first server matching the operation's channel.
              message:
                type: string
                title: WithAsyncAPIMessage
                description: The name of the message to use. If not set, defaults to the first message defined by the operation.
              binding:
                type: string
                title: WithAsyncAPIBinding
                description: The name of the binding to use. If not set, defaults to the first binding defined by the operation.
              payload:
                type: object
                title: WithAsyncAPIPayload
                description: The payload to call the AsyncAPI operation with, if any.
              authentication:
                $ref: '#/$defs/referenceableAuthenticationPolicy'
                title: WithAsyncAPIAuthentication
                description: The authentication policy, if any, to use when calling the AsyncAPI operation.
            required: [ document, operationRef ]
            unevaluatedProperties: false
      - title: CallGRPC
        description: Defines the GRPC call to perform.
        $ref: '#/$defs/taskBase'
        type: object
        unevaluatedProperties: false
        required: [ call, with ]
        properties:
          call:
            type: string
            const: grpc
          with:
            type: object
            title: GRPCArguments
            description: The GRPC call arguments.
            properties:
              proto:
                $ref: '#/$defs/externalResource'
                title: WithGRPCProto
                description: The proto resource that describes the GRPC service to call.
              service:
                type: object
                title: WithGRPCService
                unevaluatedProperties: false
                properties:
                  name:
                    type: string
                    title: WithGRPCServiceName
                    description: The name of the GRPC service to call.
                  host:
                    type: string
                    title: WithGRPCServiceHost
                    description: The hostname of the GRPC service to call.
                    pattern: ^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$
                  port:
                    type: integer
                    title: WithGRPCServicePost
                    description: The port number of the GRPC service to call.
                    minimum: 0
                    maximum: 65535
                  authentication:
                    $ref: '#/$defs/referenceableAuthenticationPolicy'
                    title: WithGRPCServiceAuthentication
                    description: The endpoint's authentication policy, if any.
                required: [ name, host ]
              method:
                type: string
                title: WithGRPCMethod
                description: The name of the method to call on the defined GRPC service.
              arguments:
                type: object
                title: WithGRPCArguments
                description: The arguments, if any, to call the method with.
                additionalProperties: true
            required: [ proto, service, method ]
            unevaluatedProperties: false
      - title: CallHTTP
        description: Defines the HTTP call to perform.
        $ref: '#/$defs/taskBase'
        type: object
        unevaluatedProperties: false
        required: [ call, with ]
        properties:
          call:
            type: string
            const: http
          with:
            type: object
            title: HTTPArguments
            description: The HTTP call arguments.
            properties:
              method:
                type: string
                title: WithHTTPMethod
                description: The HTTP method of the HTTP request to perform.
              endpoint:
                title: WithHTTPEndpoint
                description: The HTTP endpoint to send the request to.
                $ref: '#/$defs/endpoint'
              headers:
                type: object
                title: WithHTTPHeaders
                description: A name/value mapping of the headers, if any, of the HTTP request to perform.
              body:
                title: WithHTTPBody
                description: The body, if any, of the HTTP request to perform.
              query:
                type: object
                title: WithHTTPQuery
                description: A name/value mapping of the query parameters, if any, of the HTTP request to perform.
                additionalProperties: true
              output:
                type: string
                title: WithHTTPOutput
                description: The http call output format. Defaults to 'content'.
                enum: [ raw, content, response ]
            required: [ method, endpoint ]
            unevaluatedProperties: false
      - title: CallOpenAPI
        description: Defines the OpenAPI call to perform.
        $ref: '#/$defs/taskBase'
        type: object
        unevaluatedProperties: false
        required: [ call, with ]
        properties:
          call:
            type: string
            const: openapi
          with:
            type: object
            title: OpenAPIArguments
            description: The OpenAPI call arguments.
            properties:
              document:
                $ref: '#/$defs/externalResource'
                title: WithOpenAPIDocument
                description: The document that defines the OpenAPI operation to call.
              operationId:
                type: string
                title: WithOpenAPIOperation
                description: The id of the OpenAPI operation to call.
              parameters:
                type: object
                title: WithOpenAPIParameters
                description: A name/value mapping of the parameters of the OpenAPI operation to call.
                additionalProperties: true
              authentication:
                $ref: '#/$defs/referenceableAuthenticationPolicy'
                title: WithOpenAPIAuthentication
                description: The authentication policy, if any, to use when calling the OpenAPI operation.
              output:
                type: string
                enum: [ raw, content, response ]
                title: WithOpenAPIOutput
                description: The http call output format. Defaults to 'content'.
            required: [ document, operationId ]
            unevaluatedProperties: false
      - title: CallFunction
        description: Defines the function call to perform.
        $ref: '#/$defs/taskBase'
        type: object
        unevaluatedProperties: false
        required: [ call ]
        properties:
          call:
            type: string
            not:
              enum: ["asyncapi", "grpc", "http", "openapi"]
            description: The name of the function to call.
          with:
            type: object
            title: FunctionArguments
            description: A name/value mapping of the parameters, if any, to call the function with.
            additionalProperties: true
  forkTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: ForkTask
    description: Allows workflows to execute multiple tasks concurrently and optionally race them against each other, with a single possible winner, which sets the task's output.
    unevaluatedProperties: false
    required: [ fork ]
    properties:
      fork:
        type: object
        title: ForkTaskConfiguration
        description: The configuration of the branches to perform concurrently.
        unevaluatedProperties: false
        required: [ branches ]
        properties:
          branches:
            $ref: '#/$defs/taskList'
            title: ForkBranches
          compete:
            type: boolean
            title: ForkCompete
            description: Indicates whether or not the concurrent tasks are racing against each other, with a single possible winner, which sets the composite task's output.
            default: false
  doTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: DoTask
    description: Allows to execute a list of tasks in sequence.
    unevaluatedProperties: false
    required: [ do ]
    properties:
      do:
        $ref: '#/$defs/taskList'
        title: DoTaskConfiguration
        description: The configuration of the tasks to perform sequentially.
  emitTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: EmitTask
    description: Allows workflows to publish events to event brokers or messaging systems, facilitating communication and coordination between different components and services.
    required: [ emit ]
    unevaluatedProperties: false
    properties:
      emit:
        type: object
        title: EmitTaskConfiguration
        description: The configuration of an event's emission.
        unevaluatedProperties: false
        properties:
          event:
            type: object
            title: EmitEventDefinition
            description: The definition of the event to emit.
            properties:
              with:
                $ref: '#/$defs/eventProperties'
                title: EmitEventWith
                description: Defines the properties of event to emit.
                required: [ source, type ]
            additionalProperties: true
        required: [ event ]
  forTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: ForTask
    description: Allows workflows to iterate over a collection of items, executing a defined set of subtasks for each item in the collection. This task type is instrumental in handling scenarios such as batch processing, data transformation, and repetitive operations across datasets.
    required: [ for, do ]
    unevaluatedProperties: false
    properties:
      for:
        type: object
        title: ForTaskConfiguration
        description: The definition of the loop that iterates over a range of values.
        unevaluatedProperties: false
        properties:
          each:
            type: string
            title: ForEach
            description: The name of the variable used to store the current item being enumerated.
            default: item
          in:
            type: string
            title: ForIn
            description: A runtime expression used to get the collection to enumerate.
          at:
            type: string
            title: ForAt
            description: The name of the variable used to store the index of the current item being enumerated.
            default: index
        required: [ in ]
      while:
        type: string
        title: While
        description: A runtime expression that represents the condition, if any, that must be met for the iteration to continue.
      do:
        $ref: '#/$defs/taskList'
        title: ForTaskDo
  listenTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: ListenTask
    description: Provides a mechanism for workflows to await and react to external events, enabling event-driven behavior within workflow systems.
    required: [ listen ]
    unevaluatedProperties: false
    properties:
      listen:
        type: object
        title: ListenTaskConfiguration
        description: The configuration of the listener to use.
        unevaluatedProperties: false
        properties:
          to:
            $ref: '#/$defs/eventConsumptionStrategy'
            title: ListenTo
            description: Defines the event(s) to listen to.
        required: [ to ]
  raiseTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: RaiseTask
    description: Intentionally triggers and propagates errors.
    required: [ raise ]
    unevaluatedProperties: false
    properties:
      raise:
        type: object
        title: RaiseTaskConfiguration
        description: The definition of the error to raise.
        unevaluatedProperties: false
        properties:
          error:
            oneOf:
              - $ref: '#/$defs/error'
                title: RaiseErrorDefinition
                description: Defines the error to raise.
              - type: string
                title: RaiseErrorReference
                description: The name of the error to raise
        required: [ error ]
  runTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: RunTask
    description: Provides the capability to execute external containers, shell commands, scripts, or workflows.
    required: [ run ]
    unevaluatedProperties: false
    properties:
      run:
        type: object
        title: RunTaskConfiguration
        description: The configuration of the process to execute.
        unevaluatedProperties: false
        properties:
          await:
            type: boolean
            default: true
            title: AwaitProcessCompletion
            description: Whether to await the process completion before continuing.
        oneOf:
          - title: RunContainer
            description: Enables the execution of external processes encapsulated within a containerized environment.
            properties:
              container:
                type: object
                title: Container
                description: The configuration of the container to run.
                unevaluatedProperties: false
                properties:
                  image:
                    type: string
                    title: ContainerImage
                    description: The name of the container image to run.
                  command:
                    type: string
                    title: ContainerCommand
                    description: The command, if any, to execute on the container.
                  ports:
                    type: object
                    title: ContainerPorts
                    description: The container's port mappings, if any.
                  volumes:
                    type: object
                    title: ContainerVolumes
                    description: The container's volume mappings, if any.
                  environment:
                    type: object
                    title: ContainerEnvironment
                    description: A key/value mapping of the environment variables, if any, to use when running the configured process.
                required: [ image ]
            required: [ container ]
          - title: RunScript
            description: Enables the execution of custom scripts or code within a workflow, empowering workflows to perform specialized logic, data processing, or integration tasks by executing user-defined scripts written in various programming languages.
            properties:
              script:
                type: object
                title: Script
                description: The configuration of the script to run.
                unevaluatedProperties: false
                properties:
                  language:
                    type: string
                    title: ScriptLanguage
                    description: The language of the script to run.
                  arguments:
                    type: object
                    title: ScriptArguments
                    description: A key/value mapping of the arguments, if any, to use when running the configured script.
                    additionalProperties: true
                  environment:
                    type: object
                    title: ScriptEnvironment
                    description: A key/value mapping of the environment variables, if any, to use when running the configured script process.
                    additionalProperties: true
                oneOf:
                  - title: InlineScript
                    type: object
                    description: The script's code.
                    properties:
                      code:
                        type: string
                        title: InlineScriptCode
                    required: [ code ]
                  - title: ExternalScript
                    type: object
                    description: The script's resource.
                    properties:
                      source:
                        $ref: '#/$defs/externalResource'
                        title: ExternalScriptResource
                    required: [ source ]
                required: [ language ]
            required: [ script ]
          - title: RunShell
            description: Enables the execution of shell commands within a workflow, enabling workflows to interact with the underlying operating system and perform system-level operations, such as file manipulation, environment configuration, or system administration tasks.
            properties:
              shell:
                type: object
                title: Shell
                description: The configuration of the shell command to run.
                unevaluatedProperties: false
                properties:
                  command:
                    type: string
                    title: ShellCommand
                    description: The shell command to run.
                  arguments:
                    type: object
                    title: ShellArguments
                    description: A list of the arguments of the shell command to run.
                    additionalProperties: true
                  environment:
                    type: object
                    title: ShellEnvironment
                    description: A key/value mapping of the environment variables, if any, to use when running the configured process.
                    additionalProperties: true
                required: [ command ]
            required: [ shell ]
          - title: RunWorkflow
            description: Enables the invocation and execution of nested workflows within a parent workflow, facilitating modularization, reusability, and abstraction of complex logic or business processes by encapsulating them into standalone workflow units.
            properties:
              workflow:
                type: object
                title: SubflowConfiguration 
                description: The configuration of the workflow to run.
                unevaluatedProperties: false
                properties:
                  namespace:
                    type: string
                    title: SubflowNamespace
                    description: The namespace the workflow to run belongs to.
                  name:
                    type: string
                    title: SubflowName
                    description: The name of the workflow to run.
                  version:
                    type: string
                    default: latest
                    title: SubflowVersion
                    description: The version of the workflow to run. Defaults to latest.
                  input:
                    type: object
                    title: SubflowInput
                    description: The data, if any, to pass as input to the workflow to execute. The value should be validated against the target workflow's input schema, if specified.
                    additionalProperties: true
                required: [ namespace, name, version ]
            required: [ workflow ]
  setTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: SetTask
    description: A task used to set data.
    required: [ set ]
    unevaluatedProperties: false
    properties:
      set:
        type: object
        title: SetTaskConfiguration
        description: The data to set.
        minProperties: 1
        additionalProperties: true
  switchTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: SwitchTask
    description: Enables conditional branching within workflows, allowing them to dynamically select different paths based on specified conditions or criteria.
    required: [ switch ]
    unevaluatedProperties: false
    properties:
      switch:
        type: array
        title: SwitchTaskConfiguration
        description: The definition of the switch to use.
        minItems: 1
        items:
          type: object
          title: SwitchItem
          minProperties: 1
          maxProperties: 1
          additionalProperties:
            type: object
            title: SwitchCase
            description: The definition of a case within a switch task, defining a condition and corresponding tasks to execute if the condition is met.
            unevaluatedProperties: false
            required: [ then ]
            properties:
              when:
                type: string
                title: SwitchCaseCondition
                description: A runtime expression used to determine whether or not the case matches.
              then:
                $ref: '#/$defs/flowDirective'
                title: SwitchCaseOutcome
                description: The flow directive to execute when the case matches.
  tryTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: TryTask
    description: Serves as a mechanism within workflows to handle errors gracefully, potentially retrying failed tasks before proceeding with alternate ones.
    required: [ try, catch ]
    unevaluatedProperties: false
    properties:
      try:
        $ref: '#/$defs/taskList'
        title: TryTaskConfiguration
        description: The task(s) to perform.
      catch:
        type: object
        title: TryTaskCatch
        description: The object used to define the errors to catch.
        unevaluatedProperties: false
        properties:
          errors:
            type: object
            title: CatchErrors
            description: The configuration of a concept used to catch errors.
          as:
            type: string
            title: CatchAs
            description: The name of the runtime expression variable to save the error as. Defaults to 'error'.
          when:
            type: string
            title: CatchWhen
            description: A runtime expression used to determine whether or not to catch the filtered error.
          exceptWhen:
            type: string
            title: CatchExceptWhen
            description: A runtime expression used to determine whether or not to catch the filtered error.
          retry:
            oneOf:
              - $ref: '#/$defs/retryPolicy'
                title: RetryPolicyDefinition
                description: The retry policy to use, if any, when catching errors.
              - type: string
                title: RetryPolicyReference
                description: The name of the retry policy to use, if any, when catching errors.
          do:
            $ref: '#/$defs/taskList'
            title: TryTaskCatchDo
            description: The definition of the task(s) to run when catching an error.
  waitTask:
    type: object
    $ref: '#/$defs/taskBase'
    title: WaitTask
    description: Allows workflows to pause or delay their execution for a specified period of time.
    required: [ wait ]
    unevaluatedProperties: false
    properties:
      wait:
        $ref: '#/$defs/duration'
        title: WaitTaskConfiguration
        description: The amount of time to wait.
  flowDirective:
    title: FlowDirective
    description: Represents different transition options for a workflow.
    anyOf:
      - type: string
        enum: [ continue, exit, end ]
        default: continue
      - type: string
  referenceableAuthenticationPolicy:
    type: object
    title: ReferenceableAuthenticationPolicy
    description: Represents a referenceable authentication policy.
    unevaluatedProperties: false
    oneOf:
      - title: AuthenticationPolicyReference
        description: The reference of the authentication policy to use.
        properties:
          use:
            type: string
            minLength: 1
            title: ReferenceableAuthenticationPolicyName
            description: The name of the authentication policy to use.
        required: [use]
      - $ref: '#/$defs/authenticationPolicy'
  secretBasedAuthenticationPolicy:
    type: object
    title: SecretBasedAuthenticationPolicy
    description: Represents an authentication policy based on secrets.
    unevaluatedProperties: false
    properties:
      use:
        type: string
        minLength: 1
        title: SecretBasedAuthenticationPolicyName
        description: The name of the authentication policy to use.
    required: [use]
  authenticationPolicy:
    type: object
    title: AuthenticationPolicy
    description: Defines an authentication policy.
    oneOf:
    - title: BasicAuthenticationPolicy
      description: Use basic authentication.
      properties:
        basic:
          type: object
          title: BasicAuthenticationPolicyConfiguration
          description: The configuration of the basic authentication policy.
          unevaluatedProperties: false
          oneOf:
            - title: BasicAuthenticationProperties
              description: Inline configuration of the basic authentication policy.
              properties:
                username:
                  type: string
                  description: The username to use.
                password:
                  type: string
                  description: The password to use.
              required: [ username, password ]
            - $ref: '#/$defs/secretBasedAuthenticationPolicy'
              title: BasicAuthenticationPolicySecret
              description: Secret based configuration of the basic authentication policy.
      required: [ basic ]
    - title: BearerAuthenticationPolicy
      description: Use bearer authentication.
      properties:
        bearer:
          type: object
          title: BearerAuthenticationPolicyConfiguration
          description: The configuration of the bearer authentication policy.
          unevaluatedProperties: false
          oneOf:
            - title: BearerAuthenticationProperties
              description: Inline configuration of the bearer authentication policy.
              properties:
                token:
                  type: string
                  description: The bearer token to use.
              required: [ token ]
            - $ref: '#/$defs/secretBasedAuthenticationPolicy'
              title: BearerAuthenticationPolicySecret
              description: Secret based configuration of the bearer authentication policy.
      required: [ bearer ]
    - title: DigestAuthenticationPolicy
      description: Use digest authentication.
      properties:
        digest:
          type: object
          title: DigestAuthenticationPolicyConfiguration
          description: The configuration of the digest authentication policy.
          unevaluatedProperties: false
          oneOf:
            - title: DigestAuthenticationProperties
              description: Inline configuration of the digest authentication policy.
              properties:
                username:
                  type: string
                  description: The username to use.
                password:
                  type: string
                  description: The password to use.
              required: [ username, password ]
            - $ref: '#/$defs/secretBasedAuthenticationPolicy'
              title: DigestAuthenticationPolicySecret
              description: Secret based configuration of the digest authentication policy.
      required: [ digest ]
    - title: OAuth2AuthenticationPolicy
      description: Use OAuth2 authentication.
      properties:
        oauth2:
          type: object
          title: OAuth2AuthenticationPolicyConfiguration
          description: The configuration of the OAuth2 authentication policy.
          unevaluatedProperties: false
          oneOf:
            - type: object
              title: OAuth2ConnectAuthenticationProperties
              description: The inline configuration of the OAuth2 authentication policy.
              unevaluatedProperties: false
              allOf:
                - $ref: '#/$defs/oauth2AuthenticationProperties'
                - type: object
                  properties:
                    endpoints:
                      type: object
                      title: OAuth2AuthenticationPropertiesEndpoints
                      description: The endpoint configurations for OAuth2.
                      properties:
                        token:
                          type: string
                          format: uri-template
                          default: /oauth2/token
                          title: OAuth2TokenEndpoint
                          description: The relative path to the token endpoint. Defaults to '/oauth2/token\'.
                        revocation:
                          type: string
                          format: uri-template
                          default: /oauth2/revoke
                          title: OAuth2RevocationEndpoint
                          description: The relative path to the revocation endpoint. Defaults to '/oauth2/revoke'.
                        introspection:
                          type: string
                          format: uri-template
                          default: /oauth2/introspect
                          title: OAuth2IntrospectionEndpoint
                          description: The relative path to the introspection endpoint. Defaults to '/oauth2/introspect'.
            - $ref: '#/$defs/secretBasedAuthenticationPolicy'
              title: OAuth2AuthenticationPolicySecret
              description: Secret based configuration of the OAuth2 authentication policy.
      required: [ oauth2 ]
    - title: OpenIdConnectAuthenticationPolicy
      description: Use OpenIdConnect authentication.
      properties:
        oidc:
          type: object
          title: OpenIdConnectAuthenticationPolicyConfiguration
          description: The configuration of the OpenIdConnect authentication policy.
          unevaluatedProperties: false
          oneOf:
            - $ref: '#/$defs/oauth2AuthenticationProperties'
              title: OpenIdConnectAuthenticationProperties
              description: The inline configuration of the OpenIdConnect authentication policy.
              unevaluatedProperties: false
            - $ref: '#/$defs/secretBasedAuthenticationPolicy'
              title: OpenIdConnectAuthenticationPolicySecret
              description: Secret based configuration of the OpenIdConnect authentication policy.
      required: [ oidc ]
  oauth2AuthenticationProperties:
    type: object
    title: OAuth2AutenthicationData
    description: Inline configuration of the OAuth2 authentication policy.
    properties:
      authority:
        type: string
        format: uri-template
        title: OAuth2AutenthicationDataAuthority
        description: The URI that references the OAuth2 authority to use.
      grant:
        type: string
        enum: [ authorization_code, client_credentials, password, refresh_token, 'urn:ietf:params:oauth:grant-type:token-exchange']
        title: OAuth2AutenthicationDataGrant
        description: The grant type to use.
      client:
        type: object
        title: OAuth2AutenthicationDataClient
        description: The definition of an OAuth2 client.
        unevaluatedProperties: false
        properties:
          id:
            type: string
            title: ClientId
            description: The client id to use.
          secret:
            type: string
            title: ClientSecret
            description: The client secret to use, if any.
          assertion:
            type: string
            title: ClientAssertion
            description: A JWT containing a signed assertion with your application credentials.
          authentication:
            type: string
            enum: [ client_secret_basic, client_secret_post, client_secret_jwt, private_key_jwt, none ]
            default: client_secret_post
            title: ClientAuthentication
            description: The authentication method to use to authenticate the client.
      request:
        type: object
        title: OAuth2TokenRequest
        description: The configuration of an OAuth2 token request
        properties:
          encoding:
            type: string
            enum: [ 'application/x-www-form-urlencoded', 'application/json' ]
            default: 'application/x-www-form-urlencoded'
            title: Oauth2TokenRequestEncoding
      issuers:
        type: array
        title: OAuth2Issuers
        description: A list that contains that contains valid issuers that will be used to check against the issuer of generated tokens.
        items:
          type: string
      scopes:
        type: array
        title: OAuth2AutenthicationDataScopes
        description: The scopes, if any, to request the token for.
        items:
          type: string
      audiences:
        type: array
        title: OAuth2AutenthicationDataAudiences
        description: The audiences, if any, to request the token for.
        items:
          type: string
      username:
        type: string
        title: OAuth2AutenthicationDataUsername
        description: The username to use. Used only if the grant type is Password.
      password:
        type: string
        title: OAuth2AutenthicationDataPassword
        description: The password to use. Used only if the grant type is Password.
      subject:
        $ref: '#/$defs/oauth2Token'
        title: OAuth2AutenthicationDataSubject
        description: The security token that represents the identity of the party on behalf of whom the request is being made.
      actor:
        $ref: '#/$defs/oauth2Token'
        title: OAuth2AutenthicationDataActor
        description: The security token that represents the identity of the acting party.
  oauth2Token:
    type: object
    title: OAuth2TokenDefinition
    description: Represents an OAuth2 token.
    unevaluatedProperties: false
    properties:
      token:
        type: string
        title: OAuth2Token
        description: The security token to use.
      type:
        type: string
        title: OAuth2TokenType
        description: The type of the security token to use.
    required: [ token, type ]
  duration:
    oneOf:
      - type: object
        minProperties: 1
        unevaluatedProperties: false
        properties:
          days:
            type: integer
            title: DurationDays
            description: Number of days, if any.
          hours:
            type: integer
            title: DurationHours
            description: Number of days, if any.
          minutes:
            type: integer
            title: DurationMinutes
            description: Number of minutes, if any.
          seconds:
            type: integer
            title: DurationSeconds
            description: Number of seconds, if any.
          milliseconds:
            type: integer
            title: DurationMilliseconds
            description: Number of milliseconds, if any.
        title: DurationInline
        description: The inline definition of a duration.
      - type: string
        pattern: '^P(?!$)(\d+(?:\.\d+)?Y)?(\d+(?:\.\d+)?M)?(\d+(?:\.\d+)?W)?(\d+(?:\.\d+)?D)?(T(?=\d)(\d+(?:\.\d+)?H)?(\d+(?:\.\d+)?M)?(\d+(?:\.\d+)?S)?)?$'
        title: DurationExpression
        description: The ISO 8601 expression of a duration.
  error:
    type: object
    title: Error
    description: Represents an error.
    unevaluatedProperties: false
    properties:
      type:
        title: ErrorType
        description: A URI reference that identifies the error type.
        oneOf:
          - title: LiteralErrorType
            description: The literal error type.
            type: string
            format: uri-template
          - $ref: '#/$defs/runtimeExpression'
            title: ExpressionErrorType
            description: An expression based error type.
      status:
        type: integer
        title: ErrorStatus
        description: The status code generated by the origin for this occurrence of the error.
      instance:
        title: ErrorInstance
        description: A JSON Pointer used to reference the component the error originates from.
        oneOf:
          - title: LiteralErrorInstance
            description: The literal error instance.
            type: string
            format: json-pointer
          - $ref: '#/$defs/runtimeExpression'
            title: ExpressionErrorInstance
            description: An expression based error instance.
      title:
        type: string
        title: ErrorTitle
        description: A short, human-readable summary of the error.
      detail:
        type: string
        title: ErrorDetails
        description: A human-readable explanation specific to this occurrence of the error.
    required: [ type, status ]
  endpoint:
    title: Endpoint
    description: Represents an endpoint.
    oneOf:
      - $ref: '#/$defs/runtimeExpression'
      - title: LiteralEndpoint
        type: string
        format: uri-template
      - title: EndpointConfiguration
        type: object
        unevaluatedProperties: false
        properties:
          uri:
            title: EndpointUri
            description: The endpoint's URI.
            oneOf:
              - title: LiteralEndpointURI
                description: The literal endpoint's URI.
                type: string
                format: uri-template
              - $ref: '#/$defs/runtimeExpression'
                title: ExpressionEndpointURI
                description: An expression based endpoint's URI.
          authentication:
            $ref: '#/$defs/referenceableAuthenticationPolicy'
            title: EndpointAuthentication
            description: The authentication policy to use.
        required: [ uri ]
  eventProperties:
    type: object
    title: EventProperties
    description: Describes the properties of an event.
    properties:
      id:
        type: string
        title: EventId
        description: The event's unique identifier.
      source:
        title: EventSource
        description: Identifies the context in which an event happened.
        oneOf:
          - title: LiteralSource
            type: string
            format: uri-template
          - $ref: '#/$defs/runtimeExpression'
      type:
        type: string
        title: EventType
        description: This attribute contains a value describing the type of event related to the originating occurrence.
      time:
        title: EventTime
        description: When the event occured.
        oneOf:
          - title: LiteralTime
            type: string
            format: date-time
          - $ref: '#/$defs/runtimeExpression'
      subject:
        type: string
        title: EventSubject
        description: The subject of the event.
      datacontenttype:
        type: string
        title: EventDataContentType
        description: Content type of data value. This attribute enables data to carry any type of content, whereby format and encoding might differ from that of the chosen event format.
      dataschema:
        title: EventDataschema
        description: The schema describing the event format.
        oneOf:
          - title: LiteralDataSchema
            description: The literal event data schema.
            type: string
            format: uri-template
          - $ref: '#/$defs/runtimeExpression'
            title: ExpressionDataSchema
            description: An expression based event data schema.
    additionalProperties: true
  eventConsumptionStrategy:
    type: object
    title: EventConsumptionStrategy
    description: Describe the event consumption strategy to adopt.
    unevaluatedProperties: false
    oneOf:
      - title: AllEventConsumptionStrategy
        properties:
          all:
            type: array
            title: AllEventConsumptionStrategyConfiguration
            description: A list containing all the events that must be consumed.
            items:
              $ref: '#/$defs/eventFilter'
        required: [ all ]
      - title: AnyEventConsumptionStrategy
        properties:
          any:
            type: array
            title: AnyEventConsumptionStrategyConfiguration
            description: A list containing any of the events to consume.
            items:
              $ref: '#/$defs/eventFilter'
        required: [ any ]
      - title: OneEventConsumptionStrategy
        properties:
          one:
            $ref: '#/$defs/eventFilter'
            title: OneEventConsumptionStrategyConfiguration
            description: The single event to consume.
        required: [ one ]
  eventFilter:
    type: object
    title: EventFilter
    description: An event filter is a mechanism used to selectively process or handle events based on predefined criteria, such as event type, source, or specific attributes.
    unevaluatedProperties: false
    properties:
      with:
        $ref: '#/$defs/eventProperties'
        minProperties: 1
        title: WithEvent
        description: An event filter is a mechanism used to selectively process or handle events based on predefined criteria, such as event type, source, or specific attributes.
      correlate:
        type: object
        title: EventFilterCorrelate
        description: A correlation is a link between events and data, established by mapping event attributes to specific data attributes, allowing for coordinated processing or handling based on event characteristics.
        additionalProperties:
          type: object
          properties:
            from:
              type: string
              title: CorrelateFrom
              description: A runtime expression used to extract the correlation value from the filtered event.
            expect:
              type: string
              title: CorrelateExpect
              description: A constant or a runtime expression, if any, used to determine whether or not the extracted correlation value matches expectations. If not set, the first extracted value will be used as the correlation's expectation.
          required: [ from ]
    required: [ with ]
  extension:
    type: object
    title: Extension
    description: The definition of an extension.
    unevaluatedProperties: false
    properties:
      extend:
        type: string
        enum: [ call, composite, emit, for, listen, raise, run, set, switch, try, wait, all ]
        title: ExtensionTarget
        description: The type of task to extend.
      when:
        type: string
        title: ExtensionCondition
        description: A runtime expression, if any, used to determine whether or not the extension should apply in the specified context.
      before:
        $ref: '#/$defs/taskList'
        title: ExtensionDoBefore
        description: The task(s) to execute before the extended task, if any.
      after:
        $ref: '#/$defs/taskList'
        title: ExtensionDoAfter
        description: The task(s) to execute after the extended task, if any.
    required: [ extend ]
  externalResource:
    type: object
    title: ExternalResource
    description: Represents an external resource.
    unevaluatedProperties: false
    properties:
      name:
        type: string
        title: ExternalResourceName
        description: The name of the external resource, if any.
      endpoint:
        $ref: '#/$defs/endpoint'
        title: ExternalResourceEndpoint
        description: The endpoint of the external resource.
    required: [ endpoint ]
  persistence:
    type: object
    title: Persistence
    description: Persistence
    unevaluatedProperties: false
    properties:
      before:
        type: object
        title: Before
        description: Before
        unevaluatedProperties: false
        properties:
          get:
            type: string
            title: get
            description: get
          set:
            type: string
            title: get
            description: get
          delete:
            type: string
            title: get
            description: get
      after:
        type: object
        title: After
        description: After
        unevaluatedProperties: false
        properties:
          get:
            type: string
            title: get
            description: get
          set:
            type: string
            title: get
            description: get
          delete:
            type: string
            title: get
            description: get
  input:
    type: object
    title: Input
    description: Configures the input of a workflow or task.
    unevaluatedProperties: false
    properties:
      schema:
        $ref: '#/$defs/schema'
        title: InputSchema
        description: The schema used to describe and validate the input of the workflow or task.
      from:
        title: InputFrom
        description: A runtime expression, if any, used to mutate and/or filter the input of the workflow or task.
        oneOf:
          - type: string
          - type: object
  output:
    type: object
    title: Output
    description: Configures the output of a workflow or task.
    unevaluatedProperties: false
    properties:
      schema:
        $ref: '#/$defs/schema'
        title: OutputSchema
        description: The schema used to describe and validate the output of the workflow or task.
      as:
        title: OutputAs
        description: A runtime expression, if any, used to mutate and/or filter the output of the workflow or task.
        oneOf:
          - type: string
          - type: object
  export:
    type: object
    title: Export
    description: Set the content of the context. .
    unevaluatedProperties: false
    properties:
      schema:
        $ref: '#/$defs/schema'
        title: ExportSchema
        description: The schema used to describe and validate the workflow context.
      as:
        title: ExportAs
        description: A runtime expression, if any, used to export the output data to the context.
        oneOf:
          - type: string
          - type: object
  retryPolicy:
    type: object
    title: RetryPolicy
    description: Defines a retry policy.
    unevaluatedProperties: false
    properties:
      when:
        type: string
        title: RetryWhen
        description: A runtime expression, if any, used to determine whether or not to retry running the task, in a given context.
      exceptWhen:
        type: string
        title: RetryExcepWhen
        description: A runtime expression used to determine whether or not to retry running the task, in a given context.
      delay:
        $ref: '#/$defs/duration'
        title: RetryDelay
        description: The duration to wait between retry attempts.
      backoff:
        type: object
        title: RetryBackoff
        description: The retry duration backoff.
        unevaluatedProperties: false
        oneOf:
        - title: ConstantBackoff
          properties:
            constant:
              type: object
              description: The definition of the constant backoff to use, if any.
          required: [ constant ]
        - title: ExponentialBackOff
          properties:
            exponential:
              type: object
              description: The definition of the exponential backoff to use, if any.
          required: [ exponential ]
        - title: LinearBackoff
          properties:
            linear:
              type: object
              description: The definition of the linear backoff to use, if any.
          required: [ linear ]
      limit:
        type: object
        title: RetryLimit
        unevaluatedProperties: false
        properties:
          attempt:
            type: object
            title: RetryLimitAttempt
            unevaluatedProperties: false
            properties:
              count:
                type: integer
                title: RetryLimitAttemptCount
                description: The maximum amount of retry attempts, if any.
              duration:
                $ref: '#/$defs/duration'
                title: RetryLimitAttemptDuration
                description: The maximum duration for each retry attempt.
          duration:
            $ref: '#/$defs/duration'
            title: RetryLimitDuration
            description: The duration limit, if any, for all retry attempts.
        description: The retry limit, if any.
      jitter:
        type: object
        title: RetryPolicyJitter
        description: The parameters, if any, that control the randomness or variability of the delay between retry attempts.
        unevaluatedProperties: false
        properties:
          from:
            $ref: '#/$defs/duration'
            title: RetryPolicyJitterFrom
            description: The minimum duration of the jitter range.
          to:
            $ref: '#/$defs/duration'
            title: RetryPolicyJitterTo
            description: The maximum duration of the jitter range.
        required: [ from, to ]
  schema:
    type: object
    title: Schema
    description: Represents the definition of a schema.
    unevaluatedProperties: false
    properties:
      format:
        type: string
        default: json
        title: SchemaFormat
        description: The schema's format. Defaults to 'json'. The (optional) version of the format can be set using '{format}:{version}'.
    oneOf:
      - title: SchemaInline
        properties:
          document:
            description: The schema's inline definition.
        required: [ document ]
      - title: SchemaExternal
        properties:
          resource:
            $ref: '#/$defs/externalResource'
            title: SchemaExternalResource
            description: The schema's external resource.
        required: [ resource ]
  timeout:
    type: object
    title: Timeout
    description: The definition of a timeout.
    unevaluatedProperties: false
    properties:
      after:
        $ref: '#/$defs/duration'
        title: TimeoutAfter
        description: The duration after which to timeout.
    required: [ after ]
  runtimeExpression:
    type: string
    title: RuntimeExpression
    description: A runtime expression.
    pattern: "^\\s*\\$\\{.+\\}\\s*$"
`
