# Менеджер задач RestAPI
        СУБД - PostgeSQL
        (users <- users_lists -> todo_list <- list_items -> todo_item)

## Структура API

        POST/AUTH/SIGN-UP +
        POST/AUTH/SIGN-IN +

        GET/LISTS +
        GET/LISTS/{ID} +
        POST/LISTS +
        PUT/LISTS/{ID} +
        DELETE/LISTS/{ID} +
        GET/LISTS/{ID}/ITEMS +
        POST/LISTS/{ID}/ITEMS +

        PUT/ITEMS/{ID} +
        GET/ITEMS/{ID} +
        DELETE/ITEMS/{ID} +