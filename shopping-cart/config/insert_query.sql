-- DELIMITER //
-- CREATE PROCEDURE AlterTable()
-- BEGIN
-- 	IF NOT EXISTS( SELECT NULL
--             FROM INFORMATION_SCHEMA.COLUMNS
--            WHERE table_name = 'product'
--              AND table_schema = 'student_test'
--              AND column_name = 'product_details')  THEN

--   ALTER TABLE `product` ADD `product_details` varchar(255) NOT NULL;

--     END IF;
-- END //

-- DELIMITER ;