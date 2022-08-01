from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    @task
    def hello_world(self):
        path = "/api/v1/products/benchmark?benchmark=";
        self.client.get(f"{path}monolithic")
        self.client.get(f"{path}microservice")
        self.client.get(f"{path}microservice-sclae")