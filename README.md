# oapi-codegen "x-go-type" sample

## Customization

```diff
--- petstore-expanded.yaml	2020-08-21 00:49:32.000000000 +0900
+++ ../oapi-codegen/examples/petstore-expanded/petstore-expanded.yaml	2020-07-23 21:39:24.000000000 +0900
@@ -39,7 +39,7 @@
           required: false
           schema:
             type: integer
-            x-go-type: uint32
+            format: int32
       responses:
         '200':
           description: pet response
@@ -91,7 +91,7 @@
           required: true
           schema:
             type: integer
-            x-go-type: uint64
+            format: int64
       responses:
         '200':
           description: pet response
@@ -116,7 +116,7 @@
           required: true
           schema:
             type: integer
-            x-go-type: uint64
+            format: int64
       responses:
         '204':
           description: pet deleted
@@ -136,13 +136,12 @@
           properties:
             id:
               type: integer
-              x-go-type: uint64
+              format: int64
               description: Unique id of the pet
 
     NewPet:
       required:
         - name
-        - custom_date
       properties:
         name:
           type: string
@@ -150,9 +149,6 @@
         tag:
           type: string
           description: Type of the pet
-        custom_date:
-          type: string
-          x-go-type: CustomDate
 
     Error:
       required:
@@ -161,7 +157,7 @@
       properties:
         code:
           type: integer
-          x-go-type: uint32
+          format: int32
           description: Error code
         message:
           type: string

```

## How to run

1. go run server/petstore.go

