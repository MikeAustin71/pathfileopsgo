System.register(["angular2/angular2", "./todoService"], function(exports_1) {
    var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
        var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
        if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
        else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
        return c > 3 && r && Object.defineProperty(target, key, r), r;
    };
    var __metadata = (this && this.__metadata) || function (k, v) {
        if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
    };
    var angular2_1, todoService_1;
    var TodoInput;
    return {
        setters:[
            function (angular2_1_1) {
                angular2_1 = angular2_1_1;
            },
            function (todoService_1_1) {
                todoService_1 = todoService_1_1;
            }],
        execute: function() {
            TodoInput = (function () {
                function TodoInput(todoService) {
                    this.todoService = todoService;
                    console.log(todoService);
                }
                TodoInput.prototype.onClick = function (event, value) {
                    this.todoService.addTodo(value);
                    console.log(this.todoService.todos);
                };
                TodoInput = __decorate([
                    angular2_1.Component({
                        selector: 'todo-input'
                    }),
                    angular2_1.View({
                        template: "\n        <input type=\"text\" #log-me>\n        <button (click)=\"onClick($event, logMe.value)\">Log Input</button>\n    "
                    }), 
                    __metadata('design:paramtypes', [todoService_1.TodoService])
                ], TodoInput);
                return TodoInput;
            })();
            exports_1("TodoInput", TodoInput);
        }
    }
});
//# sourceMappingURL=todoInput.js.map