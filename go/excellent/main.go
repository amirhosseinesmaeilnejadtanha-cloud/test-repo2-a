package main
func EvenOrOdd(number int) string {
    if number%2 == 0 {
        return "even"
    } else {
        return "odd"
    }
}
name: platform_context
description: The context for GitOps platform, this will drive GitOps specific policies
owner: 
resource: repository
where: 
configuration:
  platformContext:
    active: true
onFailure: 
onSuccess: