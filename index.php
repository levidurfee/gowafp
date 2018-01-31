<!doctype html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <title>Hello, world!</title>
    <style media="screen">
        body {
            color:#424242;
        }
        .container {
            max-width:300px;
            margin:0 auto;
            font-family:"Courier New", Courier, monospace;
        }
        label {
            display:block;
        }
        textarea {
            width:100%;
        }
        button {
            background:#424242;
            border:0;
            color:#fff;
            padding:10px;
        }
        textarea, button {
            border-radius: 5px;
        }
        button:hover {
            cursor:pointer;
        }
        code {
            color:#fff;
            background:#E91E63;
            padding:2px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Hello, world!</h1>
        <form action="" method="POST">
            <label for="secret">Sensitive Data</label>
            <textarea class="form-control" id="secret" rows="3" name="secret"></textarea>
            <p>Type <code>hack</code> in the textbox and submit it.</p>
            <button type="submit" class="btn btn-primary mb-2">Hack!</button>
            <input type="hidden" name="hidden" value="can't see me">
        </form>
        <?php
        echo "<h2>Request Method</h2><pre>"; var_dump($_SERVER["REQUEST_METHOD"]); echo "</pre>";
        echo "<h2>Get</h2><pre>"; var_dump($_GET); echo "</pre>";
        echo "<h2>Post</h2><pre>"; var_dump($_POST); echo "</pre>";
        ?>
        <p>
            <small>
                <?php echo PHP_VERSION; ?>
            </small>
        </p>
    </div>
</body>
</html>
