> docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
> rabbitmqctl status
> rabbitmqctl list_queues
> rabbitmqctl cluster_status
> https://www.rabbitmq.com/rabbitmqctl.8.html
> rabbitmq-plugins list
> rabbitmq-plugins diable rabbitmq_mamagement
> https://www.rabbitmq.com/rabbitmq-plugins.8.html
> localhost:15672 guest/guest
> rabbitmqctl   add_user username password
> rabbitmqctl   set_user_tags username  permission_tag(administrator)
> rabbitmqctl   set_permissions -p / username  ".*" ".*" ".*" // read/write/configure permissions
> mnisia database
 > https://www.rabbitmq.com/configure.html
