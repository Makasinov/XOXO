var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/:lesson_id', function(req, res, next) {
  res.render('lesson', { title: `Урок ${req.params.lesson_id}` });
});

module.exports = router;
